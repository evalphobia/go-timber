package timber

import (
	"sync"
	"time"
)

// Daemon is background daemon for sending logs.
// This struct stores logs and sends logs to Timber.io in each checkpoint timing.
type Daemon struct {
	flushLogs func([]*logPayload) error

	spoolMu            sync.Mutex
	spool              []*logPayload
	checkpointSize     int
	checkpointInterval time.Duration
	stopSignal         chan struct{}
}

// NewDaemon creates new Daemon.
// size is number of logs to send Timber.io API in single checkpoint.
// interval is the time of checkpoint interval.
// fn is function called in each checkpoint, to sends logs to Timber.io API.
func NewDaemon(size int, interval time.Duration, fn func([]*logPayload) error) *Daemon {
	if size < 1 {
		size = 10
	}
	if interval == 0 {
		interval = 1 * time.Second
	}

	return &Daemon{
		spool:              make([]*logPayload, 0, 4096),
		checkpointSize:     size,
		checkpointInterval: interval,
		stopSignal:         make(chan struct{}),
		flushLogs:          fn,
	}
}

// Add adds logs data into daemon.
func (d *Daemon) Add(logs ...*logPayload) {
	d.spoolMu.Lock()
	d.spool = append(d.spool, logs...)
	d.spoolMu.Unlock()
}

// Flush gets logs from the internal spool and execute flushLogs function.
func (d *Daemon) Flush() {
	d.spoolMu.Lock()
	var logs []*logPayload
	logs, d.spool = shiftLogs(d.spool, d.checkpointSize)
	d.spoolMu.Unlock()
	d.flushLogs(logs)
}

// shiftLogs retrieves logs.
func shiftLogs(slice []*logPayload, size int) (part []*logPayload, all []*logPayload) {
	l := len(slice)
	if l <= size {
		return slice, slice[:0]
	}
	return slice[:size], slice[size:]
}

// Run sets timer to flush data in each checkpoint as a background daemon.
func (d *Daemon) Run() {
	ticker := time.NewTicker(d.checkpointInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				d.Flush()
			case <-d.stopSignal:
				ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops daemon.
func (d *Daemon) Stop() {
	d.stopSignal <- struct{}{}
}

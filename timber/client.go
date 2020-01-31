package timber

import (
	"time"

	"github.com/evalphobia/httpwrapper/request"
)

// Client handles logging to Timber.io.
type Client struct {
	Config
	daemon *Daemon

	systemContent *SystemContext
}

// New creates initialized *Client.
func New(conf Config) (*Client, error) {
	conf.Init()

	cli := &Client{
		Config: conf,
	}

	if !conf.Sync {
		cli.RunDaemon(conf.getCheckpointSize(), conf.getCheckpointInterval())
	}

	cli.systemContent = &SystemContext{
		Hostname: conf.hostname,
		PID:      conf.pid,
	}

	return cli, nil
}

// Debug sends a debug level log.
func (c *Client) Debug(line string, opt ...map[string]interface{}) error {
	return c.EmitWithLevel(LogLevelDebug, line, opt...)
}

// Trace sends a trace level log.
func (c *Client) Trace(line string, opt ...map[string]interface{}) error {
	return c.EmitWithLevel(LogLevelTrace, line, opt...)
}

// Info sends a info level log.
func (c *Client) Info(line string, opt ...map[string]interface{}) error {
	return c.EmitWithLevel(LogLevelInfo, line, opt...)
}

// Warn sends a warning level log.
func (c *Client) Warn(line string, opt ...map[string]interface{}) error {
	return c.EmitWithLevel(LogLevelWarn, line, opt...)
}

// Err sends a error level log.
func (c *Client) Err(line string, opt ...map[string]interface{}) error {
	return c.EmitWithLevel(LogLevelError, line, opt...)
}

// Fatal sends a fatal level log.
func (c *Client) Fatal(line string, opt ...map[string]interface{}) error {
	return c.EmitWithLevel(LogLevelFatal, line, opt...)
}

// Emit sends a log of a given level.
func (c *Client) EmitWithLevel(level, line string, opt ...map[string]interface{}) error {
	if !isMoreLevel(c.MinimumLevel, level) {
		return nil
	}

	d := LogData{
		Message: line,
		Level:   level,
	}
	if len(opt) != 0 {
		d.Extra = opt[0]
	}

	return c.emit(d)
}

// Emit sends a log.
func (c *Client) Emit(line string, opt ...map[string]interface{}) error {
	d := LogData{
		Message: line,
	}
	if len(opt) != 0 {
		d.Extra = opt[0]
	}

	return c.emit(d)
}

func (c *Client) emit(d LogData) error {
	if c.Sync {
		return c.send([]*logPayload{d.toPayload(c.systemContent)})
	}

	c.daemon.Add(d.toPayload(c.systemContent))
	return nil
}

// RunDaemon runs a Daemon in background.
func (c *Client) RunDaemon(size int, interval time.Duration) {
	c.daemon = NewDaemon(size, interval, c.send)
	c.daemon.Run()
}

// Send actually sends logs to Timber.io via HTTP API .
func (c *Client) send(logs []*logPayload) error {
	if len(logs) == 0 {
		return nil
	}

	data, err := logsToNDJson(logs)
	if err != nil {
		return err
	}

	return c.callAPI(string(data))
}

// callAPI sends POST request to endpoint.
func (c *Client) callAPI(params interface{}) error {
	conf := c.Config
	resp, err := request.POST(conf.url, request.Option{
		Payload:     params,
		PayloadType: request.PayloadTypeJSON,
		// ContentType: "application/msgpack",
		Bearer:    conf.apikey,
		Retry:     !conf.NoRetry,
		Debug:     conf.Debug,
		UserAgent: "go-timber/v0.0.1",
		Timeout:   conf.timeout,
	})
	if err != nil {
		return err
	}
	defer resp.Close()
	return nil
}

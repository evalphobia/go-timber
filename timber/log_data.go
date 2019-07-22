package timber

import (
	"encoding/json"
	"time"
)

type LogData struct {
	Time    time.Time
	Message string

	Level string
	Env   string
	Extra map[string]interface{}
}

func (d LogData) GetTime() time.Time {
	if d.Time.IsZero() {
		return time.Now()
	}
	return d.Time
}

func (d LogData) toPayload(sctx *SystemContext) *logPayload {
	return &logPayload{
		Dt:      d.GetTime(),
		Message: d.Message,
		Level:   d.Level,
		Extra:   d.Extra,
		Context: &ContextData{
			SystemContext: sctx,
		},
	}
}

type logPayload struct {
	Dt      time.Time    `json:"dt"`
	Level   string       `json:"level,omitempty"`
	Message string       `json:"message,omitempty"`
	Context *ContextData `json:"context,omitempty"`
	Extra   interface{}  `json:"extra,omitempty"`
}

type ContextData struct {
	*RuntimeContext `json:"runtime,omitempty"`
	*SystemContext  `json:"system,omitempty"`
	*UserContext    `json:"user,omitempty"`
}

type RuntimeContext struct {
	Function string `json:"function,omitempty"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
}

type SystemContext struct {
	Hostname string `json:"hostname,omitempty"`
	PID      int    `json:"pid,omitempty"`
}
type UserContext struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func logsToNDJson(logs []*logPayload) ([]byte, error) {
	result := make([]byte, 0, 4096)
	for i, log := range logs {
		b, err := json.Marshal(log)
		if err != nil {
			return nil, err
		}
		if i != 0 {
			result = append(result, []byte{'\n'}...)
		}
		result = append(result, b...)
	}
	return result, nil
}

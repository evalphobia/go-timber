package timber

const (
	LogLevelDebug = "debug"
	LogLevelTrace = "trace"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
	LogLevelFatal = "fatal"
)

var levels = map[string]int{
	LogLevelDebug: 1,
	LogLevelTrace: 2,
	LogLevelInfo:  3,
	LogLevelWarn:  4,
	LogLevelError: 5,
	LogLevelFatal: 6,
}

func isMoreLevel(minLevel, level string) bool {
	min, ok := levels[minLevel]
	if !ok {
		return true
	}

	cur, ok := levels[level]
	if !ok {
		return true
	}
	return cur >= min
}

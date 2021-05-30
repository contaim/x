package logx

// Level is used to configure the logs persisted.
type Level int32

const (
	// Trace is the most verbose level. Intended to be used for the tracing
	// of actions in code, such as function enters/exits, etc.
	Trace Level = iota + 1

	// Debug information for programmer lowlevel analysis.
	Debug

	// Info information about steady state operations.
	Info

	// Warn information about rare but handled events.
	Warn

	// Error information about unrecoverable events.
	Error

	// Off disables all logging output.
	Off
)

func (l Level) String() string {
	switch l {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Off:
		return "off"
	}
	return "invalid log level"
}

func LevelFromString(level string) Level {
	switch level {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn", "warning":
		return Warn
	case "error":
		return Error
	case "off", "none":
		return Off
	default:
		return Info
	}
}

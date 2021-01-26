package logger

type Logger interface {
	Debug(msg string, metadata map[string]interface{})
	Info(msg string, metadata map[string]interface{})
	Warn(msg string, metadata map[string]interface{})
	Error(msg string, metadata map[string]interface{})
}

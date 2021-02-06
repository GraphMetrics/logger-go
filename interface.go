package logger

import "github.com/graphmetrics/logger-go/options"

type Logger interface {
	Debug(msg string, metadata map[string]interface{})
	Info(msg string, metadata map[string]interface{})
	Warn(msg string, metadata map[string]interface{})
	Error(msg string, metadata map[string]interface{})
	WithOptions(...options.LoggerOption) Logger
}

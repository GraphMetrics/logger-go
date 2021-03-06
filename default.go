package logger

import (
	"log"

	"github.com/graphmetrics/logger-go/options"
)

type defaultLogger struct {
	debug bool
}

func NewDefault(debug bool) Logger {
	return &defaultLogger{debug: debug}
}

func (l *defaultLogger) Debug(msg string, metadata map[string]interface{}) {
	if l.debug {
		log.Printf("[DEBUG] %s %#v", msg, metadata)
	}
}

func (*defaultLogger) Info(msg string, metadata map[string]interface{}) {
	log.Printf("[INFO] %s %#v", msg, metadata)
}

func (*defaultLogger) Warn(msg string, metadata map[string]interface{}) {
	log.Printf("[WARN] %s %#v", msg, metadata)
}

func (*defaultLogger) Error(msg string, metadata map[string]interface{}) {
	log.Printf("[ERROR] %s %#v", msg, metadata)
}

func (*defaultLogger) WithOptions(...options.LoggerOption) Logger {
	return &defaultLogger{}
}

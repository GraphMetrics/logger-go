package logrusadapter

import (
	"github.com/graphmetrics/logger-go"
	"github.com/graphmetrics/logger-go/options"
	"github.com/sirupsen/logrus"
)

type logrusAdapter struct {
	parent *logrus.Logger
}

func New(parent *logrus.Logger) logger.Logger {
	return &logrusAdapter{parent: parent}
}

func (z *logrusAdapter) Debug(msg string, metadata map[string]interface{}) {
	entry := z.parent.WithFields(metadata)
	entry.Debug(msg)
}

func (z *logrusAdapter) Info(msg string, metadata map[string]interface{}) {
	entry := z.parent.WithFields(metadata)
	entry.Info(msg)
}

func (z *logrusAdapter) Warn(msg string, metadata map[string]interface{}) {
	entry := z.parent.WithFields(metadata)
	entry.Warn(msg)
}

func (z *logrusAdapter) Error(msg string, metadata map[string]interface{}) {
	entry := z.parent.WithFields(metadata)
	entry.Error(msg)
}

func (z *logrusAdapter) WithOptions(opts ...options.LoggerOption) logger.Logger {
	adapter := z
	for _, o := range opts {
		switch o.Parameter() {
		case "CallerSkipOffset":
			// Waiting on https://github.com/sirupsen/logrus/pull/1215 otherwise the caller is overriden anyway
			break
		case "Named":
			// Loggers don't have names
			break
		default:
			break
		}
	}
	return adapter
}

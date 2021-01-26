package zapadapter

import (
	"github.com/graphmetrics/logger-go"
	"go.uber.org/zap"
)

type zapAdapter struct {
	parent *zap.Logger
}

func New(parent *zap.Logger) logger.Logger {
	return &zapAdapter{parent: parent}
}

func (z *zapAdapter) Debug(msg string, metadata map[string]interface{}) {
	z.parent.Debug(msg, prepareFields(metadata)...)
}

func (z *zapAdapter) Info(msg string, metadata map[string]interface{}) {
	z.parent.Info(msg, prepareFields(metadata)...)
}

func (z *zapAdapter) Warn(msg string, metadata map[string]interface{}) {
	z.parent.Warn(msg, prepareFields(metadata)...)
}

func (z *zapAdapter) Error(msg string, metadata map[string]interface{}) {
	z.parent.Error(msg, prepareFields(metadata)...)
}

func prepareFields(metadata map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, len(metadata))
	i := 0
	for k, v := range metadata {
		fields[i] = zap.Any(k, v)
		i++
	}
	return fields
}

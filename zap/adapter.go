package zapadapter

import (
	"runtime"

	"github.com/graphmetrics/logger-go"
	"github.com/graphmetrics/logger-go/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapAdapter struct {
	parent           *zap.Logger
	callerSkipOffset int
}

func New(parent *zap.Logger) logger.Logger {
	return &zapAdapter{parent: parent, callerSkipOffset: 0}
}

func (z *zapAdapter) Debug(msg string, metadata map[string]interface{}) {
	if ce := z.parent.Check(zap.DebugLevel, msg); ce != nil {
		ce.Caller = z.caller()
		ce.Write(prepareFields(metadata)...)
	}
}

func (z *zapAdapter) Info(msg string, metadata map[string]interface{}) {
	if ce := z.parent.Check(zap.InfoLevel, msg); ce != nil {
		ce.Caller = z.caller()
		ce.Write(prepareFields(metadata)...)
	}
}

func (z *zapAdapter) Warn(msg string, metadata map[string]interface{}) {
	if ce := z.parent.Check(zap.WarnLevel, msg); ce != nil {
		ce.Caller = z.caller()
		ce.Write(prepareFields(metadata)...)
	}
}

func (z *zapAdapter) Error(msg string, metadata map[string]interface{}) {
	if ce := z.parent.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Caller = z.caller()
		ce.Write(prepareFields(metadata)...)
	}
}

func (z *zapAdapter) WithOptions(opts ...options.LoggerOption) logger.Logger {
	adapter := z
	for _, o := range opts {
		switch o.Parameter() {
		case "CallerSkipOffset":
			adapter = &zapAdapter{
				parent:           adapter.parent,
				callerSkipOffset: adapter.callerSkipOffset + o.Value().(int),
			}
			break
		case "Named":
			adapter = &zapAdapter{
				parent:           adapter.parent.Named(o.Value().(string)),
				callerSkipOffset: adapter.callerSkipOffset,
			}
			break
		default:
			break
		}
	}
	return adapter
}

func (z *zapAdapter) caller() zapcore.EntryCaller {
	return zapcore.NewEntryCaller(runtime.Caller(z.callerSkipOffset + 2))
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

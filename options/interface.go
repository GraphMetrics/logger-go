package options

type LoggerOption interface {
	Parameter() string
	Value() interface{}
}

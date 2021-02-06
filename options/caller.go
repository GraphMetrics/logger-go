package options

type CallerSkipOffset struct {
	Offset int
}

var _ LoggerOption = CallerSkipOffset{}

func (c CallerSkipOffset) Parameter() string {
	return "CallerSkipOffset"
}

func (c CallerSkipOffset) Value() interface{} {
	return c.Offset
}

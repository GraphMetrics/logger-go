package options

type Named struct {
	Name string
}

var _ LoggerOption = Named{}

func (c Named) Parameter() string {
	return "Named"
}

func (c Named) Value() interface{} {
	return c.Name
}

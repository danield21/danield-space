package content

const (
	Html Type = "text/html"
	Json Type = "appliction/json"
)

type Type string

func (t Type) AddCharset(c string) Type {
	return Type(string(t) + "; charset=" + c)
}

func (t Type) String() string {
	return string(t)
}
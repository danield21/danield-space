package content

const (
	//HTML represents the mime of an html file
	HTML Type = "text/html"
	//JSON represents the mime of an json file
	JSON Type = "appliction/json"
)

//Type is a alias for string that contains a few methods specific to the content-type
type Type string

//AddCharset adds the charset to the content type
func (t Type) AddCharset(c string) Type {
	return Type(string(t) + "; charset=" + c)
}

//String converts Type back to string
func (t Type) String() string {
	return string(t)
}

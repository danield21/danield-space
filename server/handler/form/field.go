package form

type FormField struct {
	Field        string
	ErrorMessage string
	Value        string
}

func NewFormField(field string, value string) *FormField {
	fld := new(FormField)
	*fld = FormField{field, "", value}
	return fld
}

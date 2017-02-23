package form

type Field struct {
	Field        string
	ErrorMessage string
	Value        string
}

func NewField(field string, value string) *Field {
	fld := new(Field)
	*fld = Field{field, "", value}
	return fld
}

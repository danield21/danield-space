package form

type Field struct {
	Field   string
	Message string
	Value   string
	Error   bool
}

func NewField(fld, vl string) *Field {
	field := new(Field)
	*field = Field{fld, "", vl, false}
	return field
}

func NewBoolField(fld string, vl bool) *Field {
	if vl {
		return NewField(fld, "true")
	}

	return NewField(fld, "")
}

func NewMessageField(fld, vl, msg string) *Field {
	field := new(Field)
	*field = Field{fld, msg, vl, false}
	return field
}

func NewErrorField(fld, vl, msg string) *Field {
	field := new(Field)
	*field = Field{fld, msg, vl, true}
	return field
}

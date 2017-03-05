package form

type Form struct {
	Fields    []*Field
	Messages  []string
	Submitted bool
	Error     bool
}

func NewForm(fields ...*Field) *Form {
	f := new(Form)
	for _, fld := range fields {
		f.AddField(fld)
	}
	return f
}

func NewSubmittedForm(fields ...*Field) *Form {
	f := NewForm(fields...)
	f.Submitted = true
	return f
}

func NewErrorForm(msg string) *Form {
	f := new(Form)
	f.AddErrorMessage(msg)
	return f
}

func (f *Form) AddMessage(msg string) {
	f.Messages = append(f.Messages, msg)
}

func (f *Form) AddErrorMessage(msg string) {
	f.Error = true
	f.Messages = append(f.Messages, msg)
}

func (f *Form) AddField(fld *Field) {
	f.Fields = append(f.Fields, fld)
}

func (f Form) HasErrors() bool {
	return f.Error || len(f.ErrorForm().Fields) > 0
}

func (f Form) IsSuccessful() bool {
	return !f.IsEmpty() && !f.HasErrors()
}

func (f Form) IsEmpty() bool {
	return len(f.Fields) == 0
}

func (f *Form) ErrorForm() *Form {
	errs := NewForm()
	for _, fld := range f.Fields {
		if fld != nil && fld.Error {
			errs.Fields = append(errs.Fields, fld)
		}
	}
	return errs
}

func (f *Form) FieldNames() []string {
	var flds []string
	for _, fld := range f.Fields {
		if fld != nil && fld.Error {
			flds = append(flds, fld.Field)
		}
	}
	return flds
}

func (f Form) Get(field string) *Field {
	for _, fld := range f.Fields {
		if fld != nil && fld.Field == field {
			return fld
		}
	}
	return NewField(field, "")
}

func (f Form) Has(field string) bool {
	for _, fld := range f.Fields {
		if fld != nil && fld.Field == field {
			return false
		}
	}

	return true
}

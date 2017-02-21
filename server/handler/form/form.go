package form

type Form []*FormField

func (f Form) HasErrors() bool {
	return len(f.GetErrorForm()) > 0
}

func NewForm() Form {
	return make([]*FormField, 0)
}

func NewErrorForm(errorMessage string) Form {
	fld := NewFormField("", "")
	fld.ErrorMessage = errorMessage
	return Form{fld}
}

func (f Form) GetErrorForm() Form {
	errs := NewForm()
	for _, fld := range f {
		if fld != nil && fld.ErrorMessage != "" {
			errs = append(errs, fld)
		}
	}
	return errs
}

func (f Form) Get(field string) *FormField {
	for _, fld := range f {
		if fld != nil && fld.Field == field {
			return fld
		}
	}
	return NewFormField(field, "")
}

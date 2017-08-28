package form

import (
	"net/url"
)

type Form struct {
	Fields    map[string]*Field
	Error     error
	Submitted bool
}

func MakeForm() Form {
	return Form{
		Fields: make(map[string]*Field),
	}
}

func NewForm() Form {
	return Form{
		Fields: make(map[string]*Field),
	}
}

func NewSubmittedForm() Form {
	return Form{
		Fields:    make(map[string]*Field),
		Submitted: true,
	}
}

func NewErrorForm(err error) Form {
	return Form{
		Fields: make(map[string]*Field),
		Error:  err,
	}
}

func (f Form) AddFieldFromValue(name string, values url.Values) *Field {
	fld := new(Field)
	fld.Values = values[name]

	f.Fields[name] = fld

	return fld
}

func (f Form) HasErrors() bool {
	if f.Error != nil {
		return true
	}

	for _, fld := range f.Fields {
		if fld.Error != nil {
			return true
		}
	}

	return false
}

func (f Form) IsSuccessful() bool {
	return f.Submitted && !f.HasErrors()
}

func (f Form) IsEmpty() bool {
	return len(f.Fields) == 0
}

func (f Form) ErrorForm() Form {
	errs := MakeForm()
	for name, fld := range f.Fields {
		if fld.Error != nil {
			errs.Fields[name] = fld
		}
	}
	return errs
}

func (f Form) FieldNames() []string {
	var flds []string
	for name := range f.Fields {
		flds = append(flds, name)
	}
	return flds
}

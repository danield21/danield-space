package form

import "golang.org/x/net/context"

type formKeyType string

const formKey = formKeyType("errorForm")

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

func WithForm(ctx context.Context, form Form) context.Context {
	return context.WithValue(ctx, formKey, form)
}

func GetForm(ctx context.Context) Form {
	iForm := ctx.Value(formKey)
	if iForm == nil {
		return NewForm()
	}
	form, ok := iForm.(Form)
	if !ok {
		return NewForm()
	}
	return form
}

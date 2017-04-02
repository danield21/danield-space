package form

import "golang.org/x/net/context"

type formKeyType string

const formKey = formKeyType("form")

func WithForm(ctx context.Context, form *Form) context.Context {
	return context.WithValue(ctx, formKey, form)
}

func AsForm(ctx context.Context) *Form {
	iForm := ctx.Value(formKey)
	if iForm == nil {
		return NewForm()
	}
	form, ok := iForm.(*Form)
	if !ok {
		return NewForm()
	}
	return form
}

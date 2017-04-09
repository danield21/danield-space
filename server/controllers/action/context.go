package action

import (
	"github.com/danield21/danield-space/server/form"
	"golang.org/x/net/context"
)

type formKeyType string

const formKey = formKeyType("form")

func WithForm(ctx context.Context, form form.Form) context.Context {
	return context.WithValue(ctx, formKey, form)
}

func Form(ctx context.Context) form.Form {
	iForm := ctx.Value(formKey)
	if iForm == nil {
		return form.MakeForm()
	}
	frm, ok := iForm.(form.Form)
	if !ok {
		return form.MakeForm()
	}
	return frm
}

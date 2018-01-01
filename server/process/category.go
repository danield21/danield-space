package process

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/sessions"
)

const catTitleKey = "title"
const catURLKey = "url"
const catDscKey = "description"

func UnpackCategory(values url.Values) (*store.Category, form.Form) {
	frm := form.NewSubmittedForm()

	ttlFld := frm.AddFieldFromValue(catTitleKey, values)
	form.NotEmpty(ttlFld, "Title is required")

	urlFld := frm.AddFieldFromValue(catURLKey, values)
	if !form.NotEmpty(urlFld, "URL is required") && !store.ValidURLPart(urlFld.Get()) {
		form.Fail(urlFld, "url is not in a proper format")
	}

	dscFld := frm.AddFieldFromValue(catDscKey, values)
	form.NotEmpty(dscFld, "Description is required")

	if frm.HasErrors() {
		return nil, frm
	}

	category := new(store.Category)
	*category = store.Category{
		Title:       ttlFld.Get(),
		URL:         urlFld.Get(),
		Description: dscFld.Get(),
	}
	return category, frm
}

type PutCategoryProcessor struct {
	Category store.CategoryRepository
}

func (prc PutCategoryProcessor) Process(ctx context.Context, req *http.Request, ses *sessions.Session) form.Form {
	err := req.ParseForm()
	if err != nil {
		return form.NewErrorForm(errors.New("Unable to parse form"))
	}

	cat, frm := UnpackCategory(req.Form)
	if cat == nil {
		return frm
	}

	err = prc.Category.Set(ctx, cat)
	if err != nil {
		frm.Error = errors.New("Unable to put into database")
	}

	return frm
}

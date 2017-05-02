package action

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/categories"
	"golang.org/x/net/context"
)

const catTitleKey = "title"
const catURLKey = "url"
const catDscKey = "description"

func UnpackCategory(values url.Values) (*categories.Category, form.Form) {
	frm := form.MakeForm()
	frm.Submitted = true

	ttlFld := frm.AddFieldFromValue(catTitleKey, values)
	form.NotEmpty(ttlFld, "Title is required")

	urlFld := frm.AddFieldFromValue(catURLKey, values)
	if !form.NotEmpty(urlFld, "URL is required") && !repository.ValidURLPart(urlFld.Get()) {
		form.Fail(urlFld, "url is not in a proper format")
	}

	dscFld := frm.AddFieldFromValue(catDscKey, values)
	form.NotEmpty(dscFld, "Description is required")

	if frm.HasErrors() {
		return nil, frm
	}

	category := new(categories.Category)
	*category = categories.Category{
		Title:       ttlFld.Get(),
		URL:         urlFld.Get(),
		Description: dscFld.Get(),
	}
	return category, frm
}

func PutCategoryLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, form.Form{Error: errors.New("Unable to parse form")}), e, w)
		}

		cat, frm := UnpackCategory(r.Form)
		if cat == nil {
			return h(WithForm(ctx, frm), e, w)
		}

		err = categories.Set(ctx, cat)
		if err != nil {
			frm.Error = errors.New("Unable to put into database")
		}

		return h(WithForm(ctx, frm), e, w)
	}
}

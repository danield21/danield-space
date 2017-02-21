package form

import (
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/service"
	"golang.org/x/net/context"
)

const catTitleKey = "title"
const catURLKey = "url"
const catDscKey = "description"

func UnpackCategory(values url.Values) (*categories.Category, Form) {
	title := NewFormField(catTitleKey, values.Get(catTitleKey))
	NotEmpty(title, "Title is required")

	urlFld := NewFormField(catURLKey, values.Get(catURLKey))
	if !NotEmpty(urlFld, "URL is required") && !repository.ValidUrlPart(urlFld.Value) {
		urlFld.ErrorMessage = "URL is not in a proper format"
	}

	description := NewFormField(catDscKey, values.Get(catDscKey))
	NotEmpty(description, "Description is required")

	form := Form{title, urlFld, description}

	if form.HasErrors() {
		return nil, form
	}

	category := new(categories.Category)
	*category = categories.Category{
		Title:       title.Value,
		Url:         urlFld.Value,
		Description: description.Value,
	}
	return category, form
}

func PutCategoryLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		r := service.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, NewErrorForm("Unable to parse form")), e, w)
		}

		cat, form := UnpackCategory(r.Form)
		if cat == nil {
			return h(WithForm(ctx, form), e, w)
		}

		err = categories.Set(ctx, cat)
		if err != nil {
			return h(WithForm(ctx, NewErrorForm("Unable to put into database")), e, w)
		}

		return h(WithForm(ctx, form), e, w)
	}
}

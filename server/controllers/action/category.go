package action

import (
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/categories"
	"golang.org/x/net/context"
)

const catTitleKey = "title"
const catURLKey = "url"
const catDscKey = "description"

func UnpackCategory(values url.Values) (*categories.Category, form.Form) {
	title := form.NewField(catTitleKey, values.Get(catTitleKey))
	form.NotEmpty(title, "Title is required")

	urlFld := form.NewField(catURLKey, values.Get(catURLKey))
	if !form.NotEmpty(urlFld, "URL is required") && !repository.ValidURLPart(urlFld.Value) {
		urlFld.ErrorMessage = "URL is not in a proper format"
	}

	description := form.NewField(catDscKey, values.Get(catDscKey))
	form.NotEmpty(description, "Description is required")

	f := form.Form{title, urlFld, description}

	if f.HasErrors() {
		return nil, f
	}

	category := new(categories.Category)
	*category = categories.Category{
		Title:       title.Value,
		URL:         urlFld.Value,
		Description: description.Value,
	}
	return category, f
}

func PutCategoryLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(form.WithForm(ctx, form.NewErrorForm("Unable to parse form")), e, w)
		}

		cat, f := UnpackCategory(r.Form)
		if cat == nil {
			return h(form.WithForm(ctx, f), e, w)
		}

		err = categories.Set(ctx, cat)
		if err != nil {
			errField := form.NewField("", "")
			errField.ErrorMessage = "Unable to put into database"

			f = append(f, errField)
			return h(form.WithForm(ctx, f), e, w)
		}

		return h(form.WithForm(ctx, f), e, w)
	}
}

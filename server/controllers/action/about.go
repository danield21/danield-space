package action

import (
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/repository/about"
	"golang.org/x/net/context"
)

const aboutKey = "about"

func UnpackAbout(values url.Values) ([]byte, form.Form) {
	abt := form.NewField(aboutKey, values.Get(aboutKey))
	form.NotEmpty(abt, "Title is required")

	f := form.Form{abt}

	return []byte(abt.Value), f
}

func PutAboutLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(form.WithForm(ctx, form.NewErrorForm("Unable to parse form")), e, w)
		}

		abt, f := UnpackAbout(r.Form)
		if abt == nil {
			return h(form.WithForm(ctx, f), e, w)
		}

		err = about.Set(ctx, abt)
		if err != nil {
			errField := form.NewField("", "")
			errField.ErrorMessage = "Unable to put into database"

			f = append(f, errField)
			return h(form.WithForm(ctx, f), e, w)
		}

		return h(form.WithForm(ctx, f), e, w)
	}
}

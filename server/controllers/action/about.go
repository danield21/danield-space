package action

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/about"
	"golang.org/x/net/context"
)

const aboutKey = "about"

func UnpackAbout(values url.Values) ([]byte, form.Form) {
	frm := form.MakeForm()

	fld := frm.AddFieldFromValue(aboutKey, values)
	form.NotEmpty(fld, "About is required")

	frm.Submitted = true

	return []byte(fld.Get()), frm
}

func PutAboutLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, form.Form{Error: errors.New("Unable to parse form")}), e, w)
		}

		abt, f := UnpackAbout(r.Form)
		if f.HasErrors() {
			return h(WithForm(ctx, f), e, w)
		}

		err = about.Set(ctx, abt)
		if err != nil {
			f.Error = errors.New("Unable to put into database")
		}

		return h(WithForm(ctx, f), e, w)
	}
}

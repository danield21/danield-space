package process

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

const aboutKey = "about"

func UnpackAbout(values url.Values) ([]byte, form.Form) {
	frm := form.NewSubmittedForm()

	fld := frm.AddFieldFromValue(aboutKey, values)
	form.NotEmpty(fld, "About is required")

	return []byte(fld.Get()), frm
}

type PutAboutProcessor struct {
	About store.AboutRepository
}

func (prc PutAboutProcessor) Process(ctx context.Context, r *http.Request, s *sessions.Session) form.Form {
	err := r.ParseForm()
	if err != nil {
		frm := form.NewForm()
		frm.Error = errors.New("Unable to parse form")
		return frm
	}

	abt, frm := UnpackAbout(r.Form)
	if frm.HasErrors() {
		return frm
	}

	err = prc.About.Set(ctx, abt)
	if err != nil {
		log.Errorf(ctx, "Unable to put into database %v", err)
		frm.Error = errors.New("Unable to put into database")
	}

	return frm
}

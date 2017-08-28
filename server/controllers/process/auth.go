package process

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/sessions"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

const authUsrKey = "username"
const authPwdKey = "password"

var (
	ErrUnableAuth = errors.New("Unable to authenicate")
)

func UnpackAuth(values url.Values) (string, []byte, form.Form) {
	frm := form.NewSubmittedForm()

	username := frm.AddFieldFromValue(authUsrKey, values)
	form.NotEmpty(username, "Username is required")

	password := frm.AddFieldFromValue(authPwdKey, values)
	form.NotEmpty(password, "Password is required")

	return username.Get(), []byte(password.Get()), frm
}

type SignInProcessor struct {
	Account store.AccountRepository
}

func (prc SignInProcessor) Process(ctx context.Context, r *http.Request, s *sessions.Session) form.Form {
	err := r.ParseForm()

	if err != nil {
		return form.Form{Error: errors.New("Unable to parse form")}
	}

	username, password, frm := UnpackAuth(r.Form)
	if frm.HasErrors() {
		return frm
	}

	if !prc.Account.CanLogIn(ctx, username, password) {
		log.Infof(ctx, "%s attempted to login with incorrect password", username)
		frm.Error = ErrUnableAuth
		return frm
	}

	link.SetUser(s, username)

	return frm
}

func UnauthenicateLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		s := handler.Session(ctx)
		for key := range s.Values {
			delete(s.Values, key)
		}
		return h(ctx, e, w)
	}
}

var SignOutProcessor handler.ProcessorFunc = func(ctx context.Context, r *http.Request, s *sessions.Session) form.Form {
	for key := range s.Values {
		delete(s.Values, key)
	}

	return form.NewSubmittedForm()
}

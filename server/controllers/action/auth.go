package action

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/account"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

const authUsrKey = "username"
const authPwdKey = "password"

func UnpackAuth(values url.Values) (string, []byte, form.Form) {
	frm := form.MakeForm()
	frm.Submitted = true

	username := frm.AddFieldFromValue(authUsrKey, values)
	form.NotEmpty(username, "Username is required")

	password := frm.AddFieldFromValue(authPwdKey, values)
	form.NotEmpty(password, "Password is required")

	return username.Get(), []byte(password.Get()), frm
}

func AuthenicateLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		s := handler.Session(ctx)

		err := r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, form.Form{Error: errors.New("Unable to parse form")}), e, w)
		}

		username, password, frm := UnpackAuth(r.Form)
		if frm.HasErrors() {
			return h(WithForm(ctx, frm), e, w)
		}

		if !account.CanLogIn(ctx, username, password) {
			log.Infof(ctx, "%s attempted to login with incorrect password", username)
			frm.Error = errors.New("Unable to authenicate")
			return h(WithForm(ctx, frm), e, w)
		}

		link.SetUser(s, username)

		return h(WithForm(ctx, frm), e, w)
	}
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

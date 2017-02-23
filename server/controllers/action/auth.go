package action

import (
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/repository/account"
	"golang.org/x/net/context"
)

const authUsrKey = "username"
const authPwdKey = "password"

func UnpackAuth(values url.Values) (string, []byte, form.Form) {
	username := form.NewField(authUsrKey, values.Get(authUsrKey))
	form.NotEmpty(username, "Username is required")

	password := form.NewField(authPwdKey, values.Get(authPwdKey))
	form.NotEmpty(password, "Password is required")

	return username.Value, []byte(password.Value), form.Form{username, password}
}

func AuthenicateLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		s := handler.Session(ctx)

		err := r.ParseForm()
		if err != nil {
			return h(form.WithForm(ctx, form.NewErrorForm("Unable to process request")), e, w)
		}

		username, password, f := UnpackAuth(r.Form)

		if !account.IsAdmin(ctx, username, password) {
			errField := form.NewField("", "")
			errField.ErrorMessage = "Unable to authenicate"

			f = append(f, errField)
			return h(form.WithForm(ctx, f), e, w)
		}

		link.SetUser(s, username)

		return h(form.WithForm(ctx, f), e, w)
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

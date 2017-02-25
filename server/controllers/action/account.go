package action

import (
	"net/http"
	"net/url"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/repository/account"
	"golang.org/x/net/context"
)

const acctUsrKey = "username"
const acctPwdKey = "password"
const acctCfmPwdKey = "confirm-password"
const acctSprKey = "super"

func UnpackAccount(values url.Values) (*account.Account, form.Form) {
	usrFld := form.NewField(acctUsrKey, values.Get(acctUsrKey))
	if !form.NotEmpty(usrFld, "username is required") && !account.ValidUsername(usrFld.Value) {
		usrFld.ErrorMessage = "username is not in a proper format"
	}

	pwdFld := form.NewField(acctPwdKey, values.Get(acctPwdKey))
	form.NotEmpty(pwdFld, "password is required")

	cfmPwdFld := form.NewField(acctCfmPwdKey, values.Get(acctCfmPwdKey))
	form.NotEmpty(pwdFld, "confirm password is required")

	if pwdFld.Value != cfmPwdFld.Value {
		pwdFld.ErrorMessage = "passwords do not match"
	}

	sprFld := form.NewField(acctSprKey, values.Get(acctSprKey))

	f := form.Form{usrFld, pwdFld, cfmPwdFld, sprFld}

	if f.HasErrors() {
		return nil, f
	}

	acct := new(account.Account)
	*acct = account.Account{
		Username: usrFld.Value,
		Super:    sprFld.Value != "",
	}

	acct.Password([]byte(pwdFld.Value))
	return acct, f
}

func PutAccountLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		ses := handler.Session(ctx)

		err := r.ParseForm()
		if err != nil {
			return h(form.WithForm(ctx, form.NewErrorForm("Unable to parse form")), e, w)
		}

		acct, f := UnpackAccount(r.Form)
		if acct == nil {
			return h(form.WithForm(ctx, f), e, w)
		}

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		current, err := account.Get(ctx, user)
		if err != nil {
			log.Warningf(ctx, "PutAccountLink - Unable to verify account %s\n%v", user, err)
			return ctx, status.ErrUnauthorized
		}

		if !current.Super || current.Username != user {
			log.Warningf(ctx, "PutAccountLink - %s does not have access\n%v", user, err)
			return ctx, status.ErrUnauthorized
		}

		if !current.Super {
			acct.Super = false
		}

		err = account.Put(ctx, acct)
		if err != nil {
			errField := form.NewField("", "")
			errField.ErrorMessage = "Unable to put into database"

			f = append(f, errField)
		}

		return h(form.WithForm(ctx, f), e, w)
	}
}

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
const acctCfmPwdKey = "password-confirm"
const acctSprKey = "super"

func UnpackAccount(values url.Values) (*account.Account, *form.Form) {
	usrFld := form.NewField(acctUsrKey, values.Get(acctUsrKey))
	if !form.NotEmpty(usrFld, "username is required") && !account.ValidUsername(usrFld.Value) {
		form.Fail(usrFld, "username is not in a proper format")
	}

	pwdFld := form.NewField(acctPwdKey, values.Get(acctPwdKey))
	form.NotEmpty(pwdFld, "password is required")

	cfmPwdFld := form.NewField(acctCfmPwdKey, values.Get(acctCfmPwdKey))
	form.NotEmpty(pwdFld, "confirm password is required")

	if pwdFld.Value != cfmPwdFld.Value {
		form.Fail(pwdFld, "passwords do not match")
	}

	sprFld := form.NewField(acctSprKey, values.Get(acctSprKey))

	f := form.NewSubmittedForm(usrFld, pwdFld, cfmPwdFld, sprFld)

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

func AccountToForm(acct *account.Account) *form.Form {
	usrFld := form.NewField(acctUsrKey, acct.Username)
	sprFld := form.NewBoolField(acctSprKey, acct.Super)

	return form.NewForm(usrFld, sprFld)
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
		if f.HasErrors() {
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
			f.AddErrorMessage("You do not have the the ability to do this")

			return h(form.WithForm(ctx, f), e, w)
		}

		if !current.Super {
			acct.Super = false
		}

		err = account.Put(ctx, acct)
		if err != nil {
			log.Warningf(ctx, "PutAccountLink - Unable to put into database\n%v", err)
			f.AddErrorMessage("Unable to put into database")
		}

		return h(form.WithForm(ctx, f), e, w)
	}
}

package action

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

const acctUsrKey = "username"
const acctPwdKey = "password"
const acctCfmPwdKey = "passwordConfirm"
const acctSprKey = "super"

func UnpackAccount(values url.Values) (*models.Account, form.Form) {
	frm := form.MakeForm()
	frm.Submitted = true

	usrFld := frm.AddFieldFromValue(acctUsrKey, values)
	if !form.NotEmpty(usrFld, "username is required") && !models.ValidUsername(usrFld.Get()) {
		form.Fail(usrFld, "username is not in a proper format")
	}

	pwdFld := frm.AddFieldFromValue(acctPwdKey, values)
	form.NotEmpty(pwdFld, "password is required")

	cfmPwdFld := frm.AddFieldFromValue(acctCfmPwdKey, values)
	form.NotEmpty(cfmPwdFld, "confirm password is required")

	if pwdFld.Error == nil && cfmPwdFld.Error == nil && pwdFld.Get() != cfmPwdFld.Get() {
		form.Fail(pwdFld, "passwords do not match")
		form.Fail(cfmPwdFld, "passwords do not match")
	}

	sprFld := frm.AddFieldFromValue(acctSprKey, values)

	if frm.HasErrors() {
		return nil, frm
	}

	acct := new(models.Account)
	*acct = models.Account{
		Username: usrFld.Get(),
		Super:    sprFld.Get() != "",
	}

	acct.Password([]byte(pwdFld.Get()))
	return acct, frm
}

func AccountToForm(acct *models.Account) form.Form {
	frm := form.MakeForm()
	usrFld := new(form.Field)
	usrFld.Values = []string{acct.Username}

	sprFld := new(form.Field)
	if acct.Super {
		sprFld.Values = []string{"true"}
	} else {
		sprFld.Values = []string{""}
	}

	frm.Fields[acctUsrKey] = usrFld
	frm.Fields[acctSprKey] = sprFld

	return frm
}

func PutAccountLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		current, err := e.Repository().Account().Get(ctx, user)
		if err != nil {
			log.Warningf(ctx, "PutAccountLink - Unable to verify account %s\n%v", user, err)
			return ctx, status.ErrUnauthorized
		}

		err = r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, form.Form{Error: errors.New("Unable to parse form")}), e, w)
		}

		acct, frm := UnpackAccount(r.Form)
		if frm.HasErrors() {
			return h(WithForm(ctx, frm), e, w)
		}

		if !current.Super || current.Username != user {
			log.Warningf(ctx, "PutAccountLink - %s does not have access\n%v", user, err)
			frm.Error = errors.New("You do not have the the ability to do this")

			return h(WithForm(ctx, frm), e, w)
		}

		if !current.Super {
			acct.Super = false
		}

		err = e.Repository().Account().Put(ctx, acct)
		if err != nil {
			log.Warningf(ctx, "PutAccountLink - Unable to put into database\n%v", err)
			frm.Error = errors.New("Unable to put into database")
		}

		return h(WithForm(ctx, frm), e, w)
	}
}

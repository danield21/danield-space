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

const acctUsrKey = "username"
const acctPwdKey = "password"
const acctCfmPwdKey = "passwordConfirm"
const acctSprKey = "super"

func UnpackAccount(values url.Values) (*store.Account, form.Form) {
	frm := form.NewSubmittedForm()

	usrFld := frm.AddFieldFromValue(acctUsrKey, values)
	if !form.NotEmpty(usrFld, "username is required") && !store.ValidUsername(usrFld.Get()) {
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

	acct := new(store.Account)
	*acct = store.Account{
		Username: usrFld.Get(),
		Super:    sprFld.Get() != "",
	}

	acct.Password([]byte(pwdFld.Get()))
	return acct, frm
}

func AccountToForm(acct *store.Account) form.Form {
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

type PutAccountProcessor struct {
	Account store.AccountRepository
}

func (prc PutAccountProcessor) Process(ctx context.Context, req *http.Request, ses *sessions.Session) form.Form {
	user, signedIn := User(ses)
	if !signedIn {
		return form.NewErrorForm(errors.New("User is not logged in"))
	}

	current, err := prc.Account.Get(ctx, user)
	if err != nil {
		log.Warningf(ctx, "PutAccountLink - Unable to verify account %s\n%v", user, err)
		return form.NewErrorForm(errors.New("User is not logged in"))
	}

	err = req.ParseForm()
	if err != nil {
		return form.NewErrorForm(errors.New("Unable to parse form"))
	}

	acct, frm := UnpackAccount(req.Form)
	if frm.HasErrors() {
		return frm
	}

	if !current.Super || current.Username != user {
		log.Warningf(ctx, "PutAccountLink - %s does not have access\n%v", user, err)
		frm.Error = errors.New("You do not have the the ability to do this")

		return frm
	}

	if !current.Super {
		acct.Super = false
	}

	err = prc.Account.Put(ctx, acct)
	if err != nil {
		log.Warningf(ctx, "PutAccountLink - Unable to put into database\n%v", err)
		frm.Error = errors.New("Unable to put into database")
	}

	return frm
}

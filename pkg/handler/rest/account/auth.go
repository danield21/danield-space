package account

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/account"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler/rest"
	"google.golang.org/appengine/log"
)

const unlimited = -1

//Auth checks if user has correct credentials and gives them a token
func Auth(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	var (
		username string
		password []byte
	)
	ctx := e.Context(r)
	redirect := rest.GetRedirect(r)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(ctx, "account.Auth - Unable to parse form\n%v", err)
	} else {
		us, ok := r.PostForm["username"]
		if ok {
			username = us[0]
		}
		ps, ok := r.PostForm["password"]
		if ok {
			password = []byte(ps[0])
		}
	}

	isAdmin := account.IsAdmin(ctx, username, password)
	if !isAdmin {
		log.Infof(ctx, "account.Auth - Bad credentials")
		if redirect != "" {
			http.Redirect(w, r, "/admin/signin?error=no_match", http.StatusFound)
		} else {
			http.Redirect(w, r, "/admin/signin?error=no_match&redirect="+redirect, http.StatusFound)
		}
		return
	}

	log.Infof(ctx, "account.Auth - %s logged in", username)
	session := e.Session(r)
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		log.Errorf(ctx, "account.Auth - Unable to save new session\n%v", err)
	}
	if redirect == "" {
		w.Header().Set("Location", "/admin/")
		w.WriteHeader(http.StatusFound)
	} else {
		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}
}

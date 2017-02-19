package account

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/rest"
	"github.com/danield21/danield-space/server/envir"
	"google.golang.org/appengine/log"
)

func Unauth(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	session := scp.Session()
	ctx := e.Context(r)
	redirect := rest.GetRedirect(r)

	user, signedIn := session.Values["user"]
	if signedIn {
		log.Infof(ctx, "account.Auth - %s signed out", user)
	}

	for key := range session.Values {
		delete(session.Values, key)
	}

	err := session.Save(r, w)
	if err != nil {
		log.Errorf(ctx, "account.Auth - Unable to save new session\n%v", err)
	}

	if redirect == "" {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusFound)
	} else {
		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}
	return scp, nil
}

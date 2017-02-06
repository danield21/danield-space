package account

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler/rest"
	"google.golang.org/appengine/log"
)

func Unauth(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	session := e.Session(r)
	ctx := e.Context(r)
	redirect := rest.GetRedirect(r)

	user, signedIn := session.Values["user"]
	if signedIn {
		log.Infof(ctx, "account.Auth - %s signed out", user)
	}

	for key := range session.Values {
		delete(session.Values, key)
	}
	session.Save(r, w)

	if redirect == "" {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusFound)
	} else {
		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}
}

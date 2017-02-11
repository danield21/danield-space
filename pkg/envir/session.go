package envir

import (
	"net/http"

	"github.com/gorilla/sessions"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var store = sessions.NewCookieStore([]byte("3vI(DV5ytV5WuyMxU9^W5urm(B3reQtE&P*JWEV3cr6Yr32twb5^7rtE$debM)4^"))

//GetSession returns a session using a secure key
func GetSession(r *http.Request) (session *sessions.Session) {
	var err error
	session, err = store.Get(r, "danield-space")
	if err != nil {
		ctx := appengine.NewContext(r)
		log.Warningf(ctx, "envir.GetSession - Unable to get session")
	}
	return
}

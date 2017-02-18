package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/status"
)

func Authorized(h handler.Handler) handler.Handler {
	return func(e envir.Environment, w http.ResponseWriter, r *http.Request) {
		session := e.Session(r)

		_, signedIn := GetUser(session)
		if signedIn {
			h(e, w, r)
		} else {
			status.Unauthorized(e, w, r)
		}
	}
}
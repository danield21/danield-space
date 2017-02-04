package account

import (
	"encoding/json"
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/account"
	"github.com/danield21/danield-space/pkg/envir"
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
	}

	err = json.NewEncoder(w).Encode(isAdmin)
	if err != nil {
		log.Warningf(ctx, "account.Auth - Unable to encode admin into json\n%v", err)
	}
}

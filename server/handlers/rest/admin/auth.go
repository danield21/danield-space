package article

import (
	"encoding/json"
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/controllers/admin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const unlimited = -1

func Auth(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			username string
			password []byte
		)
		context := appengine.NewContext(r)

		err := r.ParseForm()
		if err != nil {
			log.Warningf(context, "admin.Auth - Unable to parse form\n%v", err)
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

		isAdmin := admin.IsAdmin(context, username, password)
		if !isAdmin {
			log.Infof(context, "admin.Auth - Bad credentials")
		}

		err = json.NewEncoder(w).Encode(isAdmin)
		if err != nil {
			log.Warningf(context, "admin.Auth - Unable to encode admin into json\n%v", err)
		}
	}
}

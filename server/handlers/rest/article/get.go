package article

import (
	"encoding/json"
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/controllers/articles"
	"github.com/danield21/danield-space/server/handlers/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const unlimited = -1

func Get(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			limit int
		)

		context := appengine.NewContext(r)

		err := r.ParseForm()
		if err != nil {
			log.Warningf(context, "article.Get - Unable to parse form\n%v", err)
		} else {
			limit = rest.GetLimit(r)
		}

		a, err := articles.GetAll(context, limit)
		if err != nil {
			log.Warningf(context, "article.Get - Unable to get articles\n%v", err)
		}

		err = json.NewEncoder(w).Encode(a)
		if err != nil {
			log.Warningf(context, "article.Get - Unable to encode articles into json\n%v", err)
		}
	}
}

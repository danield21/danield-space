package article

import (
	"encoding/json"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/rest"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"google.golang.org/appengine/log"
)

const unlimited = -1

//Get handles get requests for articles and returns a list of JSON objects
func Get(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	var (
		limit int
	)

	r := scp.Request()
	ctx := e.Context(r)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(ctx, "article.Get - Unable to parse form\n%v", err)
	} else {
		limit = rest.GetLimit(r)
	}

	a, err := articles.GetAll(ctx, limit)
	if err != nil {
		log.Warningf(ctx, "article.Get - Unable to get articles\n%v", err)
	}

	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		log.Warningf(ctx, "article.Get - Unable to encode articles into json\n%v", err)
	}
	return scp, err
}

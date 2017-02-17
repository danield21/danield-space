package category

import (
	"encoding/json"
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/categories"
	"google.golang.org/appengine/log"
)

//Get handles get requests for articles and returns a list of JSON objects
func Get(e envir.Environment, w http.ResponseWriter, r *http.Request) {

	ctx := e.Context(r)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(ctx, "article.Get - Unable to parse form\n%v", err)
	}

	c, err := categories.GetAll(ctx)
	if err != nil {
		log.Warningf(ctx, "article.Get - Unable to get articles\n%v", err)
	}

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		log.Warningf(ctx, "article.Get - Unable to encode articles into json\n%v", err)
	}
}

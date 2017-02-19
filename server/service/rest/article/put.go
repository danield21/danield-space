package article

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/service/admin"
	"github.com/danield21/danield-space/server/service/status"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"google.golang.org/appengine/log"
)

func Put(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	ctx := e.Context(r)
	session := scp.Session()
	path := mux.Vars(r)

	_, signed := admin.GetUser(session)
	if !signed {
		return scp, status.ErrUnauthorized
	}

	r.ParseForm()

	var form articles.FormArticle
	decoder := schema.NewDecoder()
	err := decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "article.Put - Unable to decode form", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return scp, err
	}

	form.Category = path["category"]
	form.Url = path["url"]

	article, err := form.Unpack(ctx)
	if err != nil {
		log.Warningf(ctx, "article.Put - Unable unpack form", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return scp, err
	}

	return scp, articles.Set(ctx, article)
}

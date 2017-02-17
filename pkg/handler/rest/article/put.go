package article

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler/admin"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"google.golang.org/appengine/log"
)

func Put(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	session := e.Session(r)
	path := mux.Vars(r)

	_, signed := admin.GetUser(session)
	if !signed {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	var form articles.FormArticle
	decoder := schema.NewDecoder()
	err := decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "article.Put - Unable to decode form", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	form.Category = path["category"]
	form.Url = path["url"]

	article, err := form.Unpack(ctx)
	if err != nil {
		log.Warningf(ctx, "article.Put - Unable unpack form", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	articles.Set(ctx, article)
}

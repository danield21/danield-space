package article

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

func Put(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := handler.Request(ctx)
	session := handler.Session(ctx)
	path := mux.Vars(r)

	_, signed := admin.GetUser(session)
	if !signed {
		return ctx, status.ErrUnauthorized
	}

	r.ParseForm()

	var form articles.FormArticle
	decoder := schema.NewDecoder()
	err := decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "article.Put - Unable to decode form", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return ctx, err
	}

	form.Category = path["category"]
	form.Url = path["url"]

	article, err := form.Unpack(ctx)
	if err != nil {
		log.Warningf(ctx, "article.Put - Unable unpack form", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return ctx, err
	}

	return ctx, articles.Set(ctx, article)
}

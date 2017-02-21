package category

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/categories"
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

	var form categories.FormCategory
	decoder := schema.NewDecoder()
	err := decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to decode form", err)
		w.WriteHeader(http.StatusBadRequest)
		return ctx, err
	}

	form.Url = path["category"]

	category, err := form.Unpack()
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable unpack form", err)
		w.WriteHeader(http.StatusBadRequest)
		return ctx, err
	}

	err = categories.Set(ctx, category)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to place category into database", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return ctx, err
}

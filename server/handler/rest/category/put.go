package category

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler/admin"
	"github.com/danield21/danield-space/server/repository/categories"
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

	var form categories.FormCategory
	decoder := schema.NewDecoder()
	err := decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to decode form", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form.Url = path["category"]

	category, err := form.Unpack()
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable unpack form", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = categories.Set(ctx, category)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to place category into database", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

package category

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/service/admin"
	"github.com/danield21/danield-space/server/service/status"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"google.golang.org/appengine/log"
)

func Put(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
	r := scp.Request()
	ctx := e.Context(r)
	session := scp.Session()
	path := mux.Vars(r)

	_, signed := admin.GetUser(session)
	if !signed {
		return status.ErrUnauthorized
	}

	r.ParseForm()

	var form categories.FormCategory
	decoder := schema.NewDecoder()
	err := decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to decode form", err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	form.Url = path["category"]

	category, err := form.Unpack()
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable unpack form", err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	err = categories.Set(ctx, category)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to place category into database", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return err
}

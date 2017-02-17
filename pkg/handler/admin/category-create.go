package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/repository/categories"
	"github.com/gorilla/schema"
	"google.golang.org/appengine/log"
)

func CategoryCreate(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(ctx, "admin.CategoryForm - Unable to parse form\n%v", err)
	}

	var form categories.FormCategory

	decoder := schema.NewDecoder()
	err = decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to decode form\n%v", err)
	}

	category, err := form.Unpack()
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable unpack form\n%v", err)
	}

	err = categories.Set(ctx, category)
	if err != nil {
		log.Warningf(ctx, "category.Put - Unable to place category into database\n%v", err)
	}

	ShowPage("category-create")(e, w, r)
}

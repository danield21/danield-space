package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/service"
	"github.com/gorilla/schema"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

func CategoryCreate(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := service.Request(ctx)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(ctx, "admin.CategoryForm - Unable to parse form\n%v", err)
		return ctx, err
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

	return ShowPage("category-create")(ctx, e, w)
}
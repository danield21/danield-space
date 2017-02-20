package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var PublicationsCategoryHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()},
)

var PublicationsCategoryPageHandler = service.Chain(
	view.HTMLHandler,
	service.ToLink(service.Chain(
		PublicationsCategoryHeadersHandler,
		PublicationsCategoryPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

//PublicationsType handles the index page
func PublicationsCategoryPageLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		r := service.Request(ctx)
		vars := mux.Vars(r)

		info := siteInfo.Get(ctx)

		cat, err := categories.Get(ctx, vars["category"])
		if err != nil {
			log.Errorf(ctx, "app.PublicationsCategoryPageLink - Unable to get category %s\n%v", vars["category"], err)
			return ctx, status.ErrNotFound
		}

		a, err := articles.GetAllByCategory(ctx, cat, 1)
		if err != nil {
			log.Errorf(ctx, "app.PublicationsCategoryPageLink - Unable to get articles by category %s\n%v", cat.Title, err)
			return ctx, status.ErrNotFound
		}

		data := struct {
			service.BaseModel
			Articles []*articles.Article
			Category *categories.Category
		}{
			BaseModel: service.BaseModel{
				SiteInfo: info,
			},
			Articles: a,
			Category: cat,
		}
		return h(link.PageContext(ctx, "page/app/publications-type", data), e, w)
	}
}

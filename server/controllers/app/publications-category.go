package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var PublicationsCategoryHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var PublicationsCategoryPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		PublicationsCategoryHeadersHandler,
		PublicationsCategoryPageLink,
		status.LinkAll,
	)),
)

//PublicationsType handles the index page
func PublicationsCategoryPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		vars := mux.Vars(r)

		info := e.Repository().SiteInfo().Get(ctx)

		cat, err := e.Repository().Category().Get(ctx, vars["category"])
		if err != nil {
			log.Errorf(ctx, "app.PublicationsCategoryPageLink - Unable to get category %s\n%v", vars["category"], err)
			return ctx, status.ErrNotFound
		}

		a, err := e.Repository().Article().GetAllByCategory(ctx, cat, 1)
		if err != nil {
			log.Errorf(ctx, "app.PublicationsCategoryPageLink - Unable to get articles by category %s\n%v", cat.Title, err)
			return ctx, status.ErrNotFound
		}

		data := struct {
			view.BaseModel
			Articles []*models.Article
			Category *models.Category
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			Articles: a,
			Category: cat,
		}
		return h(link.PageContext(ctx, "page/app/publications-type", data), e, w)
	}
}

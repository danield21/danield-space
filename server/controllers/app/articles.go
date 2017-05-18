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

var ArticleHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var ArticlePageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		ArticleHeadersHandler,
		ArticlePageLink,
		status.LinkAll,
	)),
)

func ArticlePageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		vars := mux.Vars(r)

		info := e.Repository().SiteInfo().Get(ctx)
		cat := models.NewEmptyCategory(vars["category"])

		a, err := e.Repository().Article().Get(ctx, cat, vars["key"])
		if err != nil {
			log.Errorf(ctx, "app.ArticlePageLink - Unable to get articles by type\n%v", err)
			return ctx, status.ErrNotFound
		}

		data := struct {
			view.BaseModel
			Article *models.Article
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			Article: a,
		}
		return h(link.PageContext(ctx, "page/app/article", data), e, w)
	}
}

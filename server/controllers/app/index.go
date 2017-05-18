package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"github.com/danield21/danield-space/server/repository/articles"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var IndexHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var IndexPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		IndexHeadersHandler,
		IndexPageLink,
		status.LinkAll,
	)),
)

func IndexPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := e.Repository().SiteInfo().Get(ctx)

		a, err := articles.GetAll(ctx, 10)
		if err != nil {
			log.Errorf(ctx, "app.IndexPageLink - Unable to get last 10 articles\n%v", err)
		}

		data := struct {
			view.BaseModel
			Articles []*models.Article
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			Articles: a,
		}

		return h(link.PageContext(ctx, "page/app/index", data), e, w)
	}
}

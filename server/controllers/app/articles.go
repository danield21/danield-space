package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type publicationList struct {
	Category *store.Category
	Articles []*store.Article
}

var ArticlesHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var ArticlesPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		ArticlesHeadersHandler,
		ArticlesPageLink,
		status.LinkAll,
	)),
)

//Articles handles the index page
func ArticlesPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := e.Repository().SiteInfo().Get(ctx)

		articleMap, err := e.Repository().Article().GetMapKeyedByCategory(ctx, 10)
		if err != nil {
			log.Errorf(ctx, "app.Articles - Unable to get articles organized by their type\n%v", err)
		}

		var articles []publicationList

		for cat, a := range articleMap {
			articles = append(articles, publicationList{
				Category: cat,
				Articles: a,
			})
		}

		data := struct {
			view.BaseModel
			Articles []publicationList
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			Articles: articles,
		}

		return h(link.PageContext(ctx, "page/app/articles", data), e, w)
	}
}

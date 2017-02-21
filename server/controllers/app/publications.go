package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type publicationList struct {
	Category *categories.Category
	Articles []*articles.Article
}

var PublicationsHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var PublicationsPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		PublicationsHeadersHandler,
		PublicationsPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

//Publications handles the index page
func PublicationsPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)

		articleMap, err := articles.GetMapKeyedByCategory(ctx, 10)
		if err != nil {
			log.Errorf(ctx, "app.Publications - Unable to get articles organized by their type\n%v", err)
		}

		var publications []publicationList

		for cat, a := range articleMap {
			publications = append(publications, publicationList{
				Category: cat,
				Articles: a,
			})
		}

		data := struct {
			handler.BaseModel
			Publications []publicationList
		}{
			BaseModel: handler.BaseModel{
				SiteInfo: info,
			},
			Publications: publications,
		}

		return h(link.PageContext(ctx, "page/app/publications", data), e, w)
	}
}

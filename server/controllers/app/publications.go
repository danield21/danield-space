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
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type publicationList struct {
	Category *categories.Category
	Articles []*articles.Article
}

var PublicationsHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()},
)

var PublicationsPageHandler = service.Chain(
	view.HTMLHandler,
	service.ToLink(service.Chain(
		PublicationsHeadersHandler,
		PublicationsPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

//Publications handles the index page
func PublicationsPageLink(h service.Handler) service.Handler {
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
			service.BaseModel
			Publications []publicationList
		}{
			BaseModel: service.BaseModel{
				SiteInfo: info,
			},
			Publications: publications,
		}

		return h(link.PageContext(ctx, "page/app/publications", data), e, w)
	}
}

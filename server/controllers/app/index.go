package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var IndexHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()},
)

var IndexPageHandler = service.Chain(
	view.HTMLHandler,
	service.ToLink(IndexHeadersHandler),
	IndexPageLink,
	link.Theme,
	status.LinkAll,
)

func IndexPageLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)

		a, err := articles.GetAll(ctx, 10)
		if err != nil {
			log.Errorf(ctx, "app.IndexPageLink - Unable to get last 10 articles\n%v", err)
		}

		data := struct {
			service.BaseModel
			Articles []*articles.Article
		}{
			BaseModel: service.BaseModel{
				SiteInfo: info,
			},
			Articles: a,
		}

		return h(link.PageContext(ctx, "page/app/index", data), e, w)
	}
}

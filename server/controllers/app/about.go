package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/about"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var AboutHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var AboutPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AboutHeadersHandler,
		AboutPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

func AboutPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)

		abt, err := about.Get(ctx)
		if err != nil {
			log.Errorf(ctx, "app.AboutPageLink - Unable to get About page\n%v", err)
		}

		data := struct {
			view.BaseModel
			About template.HTML
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			About: abt,
		}

		return h(link.PageContext(ctx, "page/app/about", data), e, w)
	}
}

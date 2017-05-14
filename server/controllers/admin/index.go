package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

var IndexHeadersHandler = AdminHeadersHandler

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
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)
		cats, _ := categories.GetAll(ctx)
		arts, _ := articles.GetAll(ctx, 1)

		data := struct {
			AdminModel
			HasCategories bool
			HasArticles   bool
		}{
			AdminModel: AdminModel{
				BaseModel: view.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			HasCategories: len(cats) > 0,
			HasArticles:   len(arts) > 0,
		}

		return h(link.PageContext(ctx, "page/admin/index", data), e, w)
	}
}

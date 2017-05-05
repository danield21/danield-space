package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

var CategoryHeadersHandler = AdminHeadersHandler

var CategoryPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		CategoryHeadersHandler,
		CategoryPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

func CategoryPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)
		cats, _ := categories.GetAll(ctx)

		data := struct {
			AdminModel
			Categories []*categories.Category
		}{
			AdminModel: AdminModel{
				BaseModel: view.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Categories: cats,
		}

		return h(link.PageContext(ctx, "page/admin/category", data), e, w)
	}
}

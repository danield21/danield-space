package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
)

var CategoryHeadersHandler = AdminHeadersHandler

var CategoryPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		CategoryHeadersHandler,
		CategoryPageLink,
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

		info := e.Repository().SiteInfo().Get(ctx)
		cats, _ := e.Repository().Category().GetAll(ctx)

		data := struct {
			AdminModel
			Categories []*models.Category
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

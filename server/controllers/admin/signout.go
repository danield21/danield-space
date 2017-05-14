package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

var SignOutHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var SignOutPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		SignOutHeadersHandler,
		SignOutPageLink,
		status.LinkAll,
	)),
)

var SignOutActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		SignOutHeadersHandler,
		SignOutPageLink,
		link.SaveSession,
		action.UnauthenicateLink,
		status.LinkAll,
	)),
)

func SignOutPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)

		data := struct {
			view.BaseModel
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
		}

		return h(link.PageContext(ctx, "page/admin/signout", data), e, w)
	}
}

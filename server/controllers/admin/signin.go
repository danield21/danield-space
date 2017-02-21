package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

type signinModel struct {
	handler.BaseModel
	Redirect string
}

var SignInHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var SignInPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		SignInHeadersHandler,
		SignInPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

func SignInPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		info := siteInfo.Get(ctx)

		data := signinModel{
			BaseModel: handler.BaseModel{
				SiteInfo: info,
			},
			Redirect: action.Redirect(r),
		}

		return h(link.PageContext(ctx, "page/admin/signin", data), e, w)
	}
}

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

var SignInHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var SignInPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		SignInHeadersHandler,
		SignInPageLink,
		status.LinkAll,
	)),
)

var SignInActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		SignInHeadersHandler,
		SignInPageLink,
		link.SaveSession,
		action.AuthenicateLink,
		status.LinkAll,
	)),
)

func SignInPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)
		frm := action.Form(ctx)

		result := action.Result{
			Form: frm,
		}

		data := struct {
			view.BaseModel
			action.Result
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			Result: result,
		}

		return h(link.PageContext(ctx, "page/admin/signin", data), e, w)
	}
}

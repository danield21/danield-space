package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
)

var CategoryCreateHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var CategoryCreatePageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		CategoryCreateHeadersHandler,
		CategoryCreatePageLink,
		status.LinkAll,
	)),
)

var CategoryCreateActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		CategoryCreateHeadersHandler,
		CategoryCreatePageLink,
		action.PutCategoryLink,
		status.LinkAll,
	)),
)

func CategoryCreatePageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		frm := action.Form(ctx)
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := e.Repository().SiteInfo().Get(ctx)

		data := struct {
			AdminModel
			action.Result
		}{
			AdminModel: AdminModel{
				BaseModel: view.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Result: action.Result{
				Form: frm,
			},
		}

		return h(link.PageContext(ctx, "page/admin/category-create", data), e, w)
	}
}

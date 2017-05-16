package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
)

var AdminHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

func NewAdminPageHandler(page string) handler.Handler {
	return handler.Chain(
		view.HTMLHandler,
		handler.ToLink(handler.Chain(
			AdminHeadersHandler,
			AdminPageLink(page),
			status.LinkAll,
		)),
	)
}

func AdminPageLink(page string) handler.Link {
	return func(h handler.Handler) handler.Handler {
		return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
			ses := handler.Session(ctx)

			user, signedIn := link.User(ses)

			if !signedIn {
				return ctx, status.ErrUnauthorized
			}

			info := e.Repository().SiteInfo().Get(ctx)

			data := struct {
				AdminModel
			}{
				AdminModel: AdminModel{
					BaseModel: view.BaseModel{
						SiteInfo: info,
					},
					User: user,
				},
			}

			return h(link.PageContext(ctx, page, data), e, w)
		}
	}
}

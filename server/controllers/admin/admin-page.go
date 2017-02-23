package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/siteInfo"
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
			link.Theme,
			status.LinkAll,
		)),
	)
}

func AdminPageLink(page string) handler.Link {
	return func(h handler.Handler) handler.Handler {
		return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
			session := handler.Session(ctx)

			user, signedIn := link.User(session)

			if !signedIn {
				return ctx, status.ErrUnauthorized
			}

			info := siteInfo.Get(ctx)

			data := struct {
				AdminModel
			}{
				AdminModel: AdminModel{
					BaseModel: handler.BaseModel{
						SiteInfo: info,
					},
					User: user,
				},
			}

			return h(link.PageContext(ctx, page, data), e, w)
		}
	}
}

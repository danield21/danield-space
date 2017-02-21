package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"golang.org/x/net/context"
)

var IndexHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()},
)

var IndexPageHandler = service.Chain(
	view.HTMLHandler,
	service.ToLink(service.Chain(
		IndexHeadersHandler,
		IndexPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

func IndexPageLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		session := service.Session(ctx)

		user, signedIn := GetUser(session)

		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)

		data := struct {
			AdminModel
		}{
			AdminModel: AdminModel{
				BaseModel: service.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
		}

		return h(link.PageContext(ctx, "page/admin/index", data), e, w)
	}
}

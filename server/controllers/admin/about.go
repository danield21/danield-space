package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/about"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var AboutHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var AboutPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AboutHeadersHandler,
		AboutPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

var AboutActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AboutHeadersHandler,
		AboutPageLink,
		action.PutAboutLink,
		link.Theme,
		status.LinkAll,
	)),
)

func AboutPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		f := form.AsForm(ctx)
		s := handler.Session(ctx)

		user, signedIn := link.User(s)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		if f.Has("about") {
			html, err := about.Get(ctx)
			if err == nil {
				fld := form.NewField("about", string(html))
				fld.ErrorMessage = "hidden"
				f = append(f, fld)
			} else {
				log.Warningf(ctx, "Unable to get about summary\n%v", err)
			}
		}

		info := siteInfo.Get(ctx)

		data := struct {
			AdminModel
			Form form.Form
		}{
			AdminModel: AdminModel{
				BaseModel: handler.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Form: f,
		}

		return h(link.PageContext(ctx, "page/admin/about", data), e, w)
	}
}

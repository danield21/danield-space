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
		var redirect action.URL
		f := form.AsForm(ctx)
		s := handler.Session(ctx)

		user, signedIn := link.User(s)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		if f.IsEmpty() {
			html, err := about.Get(ctx)
			if err == nil {
				fld := form.NewField("about", string(html))
				f.AddField(fld)
			} else {
				log.Warningf(ctx, "Unable to get about summary\n%v", err)
			}
		} else if f.IsSuccessful() {
			f.AddMessage("Successfully saved About summary")
			redirect = action.URL{
				URL:   "/admin/",
				Title: "Back to Admin Panel",
			}
		}

		result := action.Result{
			Form:     f,
			Redirect: redirect,
		}

		info := siteInfo.Get(ctx)

		data := struct {
			AdminModel
			action.Result
		}{
			AdminModel: AdminModel{
				BaseModel: handler.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Result: result,
		}

		return h(link.PageContext(ctx, "page/admin/about", data), e, w)
	}
}

package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
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
		status.LinkAll,
	)),
)

var AboutActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AboutHeadersHandler,
		AboutPageLink,
		action.PutAboutLink,
		status.LinkAll,
	)),
)

func AboutPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		frm := action.Form(ctx)
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		if frm.IsEmpty() {
			html, err := about.Get(ctx)
			if err == nil {
				abtFld := new(form.Field)
				abtFld.Values = []string{string(html)}
				frm.Fields["about"] = abtFld
			} else {
				log.Warningf(ctx, "Unable to get about summary\n%v", err)
			}
		}

		result := action.Result{
			Form: frm,
		}

		info := siteInfo.Get(ctx)

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
			Result: result,
		}

		return h(link.PageContext(ctx, "page/admin/about", data), e, w)
	}
}

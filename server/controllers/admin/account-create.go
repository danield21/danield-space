package admin

import (
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/account"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

var AccountCreateHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var AccountCreatePageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AccountCreateHeadersHandler,
		AccountCreatePageLink,
		link.Theme,
		status.LinkAll,
	)),
)

var AccountCreateActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AccountCreateHeadersHandler,
		AccountCreatePageLink,
		action.PutAccountLink,
		link.Theme,
		status.LinkAll,
	)),
)

func AccountCreatePageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		var redirect action.URL
		ses := handler.Session(ctx)
		req := handler.Request(ctx)
		f := form.AsForm(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		current, err := account.Get(ctx, user)
		if err != nil {
			log.Warningf(ctx, "AccountCreatePageLink - Unable to verify account %s\n%v", user, err)
			return ctx, status.ErrUnauthorized
		}

		target := req.Form.Get("account")
		if f.IsEmpty() && target != "" {
			tUser, err := account.Get(ctx, target)
			if err == nil {
				f = action.AccountToForm(tUser)
			} else {
				log.Warningf(ctx, "Unable to get account %s\n%v", target, err)
			}
		} else if f.IsSuccessful() {
			f.AddMessage("Successfully created account")
			redirect = action.URL{
				URL:   "/admin/",
				Title: "Back to Admin Panel",
			}
		}

		info := siteInfo.Get(ctx)

		data := struct {
			AdminModel
			action.Result
			Super bool
		}{
			AdminModel: AdminModel{
				BaseModel: handler.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Result: action.Result{
				Form:     f,
				Redirect: redirect,
			},
			Super: current.Super,
		}

		return h(link.PageContext(ctx, "page/admin/account-create", data), e, w)
	}
}

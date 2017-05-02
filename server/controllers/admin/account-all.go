package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/account"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var AccountAllHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var AccountAllPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		AccountAllHeadersHandler,
		AccountAllPageLink,
		link.Theme,
		status.LinkAll,
	)),
)

func AccountAllPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		current, err := account.Get(ctx, user)
		if err != nil {
			log.Warningf(ctx, "AccountAllPageLink - Unable to verify account %s\n%v", user, err)
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)
		accnts, err := account.GetAll(ctx)
		if err != nil {
			log.Debugf(ctx, "AccountAllPageLink - Unable to get all acounts\n%v", err)
		}

		data := struct {
			AdminModel
			Accounts []*account.Account
			Super    bool
		}{
			AdminModel: AdminModel{
				BaseModel: view.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Accounts: accnts,
			Super:    current.Super,
		}

		return h(link.PageContext(ctx, "page/admin/account-all", data), e, w)
	}
}

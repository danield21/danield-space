package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var PublishHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType},
)

var PublishPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		PublishHeadersHandler,
		PublishPageLink,
		status.LinkAll,
	)),
)

var PublishActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		PublishHeadersHandler,
		PublishPageLink,
		action.PutArticleLink,
		status.LinkAll,
	)),
)

//Publish handles the index page
func PublishPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		ses := handler.Session(ctx)
		frm := action.Form(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := e.Repository().SiteInfo().Get(ctx)

		cats, err := e.Repository().Category().GetAll(ctx)
		if err != nil {
			log.Warningf(ctx, "admin.Publish - Unable to get types of articles\n%v", err)
		}

		data := struct {
			AdminModel
			Categories []*store.Category
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
			Categories: cats,
		}

		return h(link.PageContext(ctx, "page/admin/publish", data), e, w)
	}
}

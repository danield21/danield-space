package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
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
		link.Theme,
		status.LinkAll,
	)),
)

var PublishActionHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		PublishHeadersHandler,
		PublishPageLink,
		action.PutArticleLink,
		link.Theme,
		status.LinkAll,
	)),
)

//Publish handles the index page
func PublishPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		var redirect action.URL
		ses := handler.Session(ctx)
		f := form.AsForm(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)

		cats, err := categories.GetAll(ctx)
		if err != nil {
			log.Warningf(ctx, "admin.Publish - Unable to get types of articles\n%v", err)
		}

		if f.IsSuccessful() {
			f.AddMessage("Successfully published article")
			redirect = action.URL{
				URL:   "/admin/",
				Title: "Back to Admin Panel",
			}
		}

		data := struct {
			AdminModel
			Categories []*categories.Category
			action.Result
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
			Categories: cats,
		}

		return h(link.PageContext(ctx, "page/admin/publish", data), e, w)
	}
}

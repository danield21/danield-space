package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/form"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"golang.org/x/net/context"
)

var CategoryCreateHeadersHandler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()},
)

var CategoryCreatePageHandler = service.Chain(
	view.HTMLHandler,
	service.ToLink(service.Chain(
		CategoryCreateHeadersHandler,
		CategoryCreatePageLink,
		link.Theme,
		status.LinkAll,
	)),
)

var CategoryCreateFormHandler = service.Chain(
	view.HTMLHandler,
	service.ToLink(service.Chain(
		CategoryCreateHeadersHandler,
		CategoryCreatePageLink,
		form.PutCategoryLink,
		link.Theme,
		status.LinkAll,
	)),
)

func CategoryCreatePageLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		f := form.GetForm(ctx)
		s := service.Session(ctx)

		user, signedIn := GetUser(s)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)

		data := struct {
			AdminModel
			Form form.Form
		}{
			AdminModel: AdminModel{
				BaseModel: service.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Form: f,
		}

		return h(link.PageContext(ctx, "page/admin/category-create", data), e, w)
	}
}

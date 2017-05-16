package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/articles"
	"golang.org/x/net/context"
)

var ArticlesHeadersHandler = AdminHeadersHandler

var ArticlesPageHandler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(handler.Chain(
		ArticlesHeadersHandler,
		ArticlePageLink,
		status.LinkAll,
	)),
)

func ArticlePageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		ses := handler.Session(ctx)

		user, signedIn := link.User(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := e.Repository().SiteInfo().Get(ctx)
		arts, _ := articles.GetAll(ctx, -1)

		data := struct {
			AdminModel
			Articles []*articles.Article
		}{
			AdminModel: AdminModel{
				BaseModel: view.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Articles: arts,
		}

		return h(link.PageContext(ctx, "page/admin/article", data), e, w)
	}
}

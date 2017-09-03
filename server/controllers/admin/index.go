package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/controllers/session"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type IndexController struct {
	Renderer            handler.Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	Category            store.CategoryRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr IndexController) Serve(ctx context.Context, pg *handler.Page, rqs *http.Request) controller.Controller {

	user, signedIn := session.User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)
	cats, _ := ctr.Category.GetAll(ctx)
	arts, _ := ctr.Article.GetAll(ctx, 1)

	cnt, err := ctr.Renderer.Render(ctx, "page/admin/index", struct {
		User          string
		HasCategories bool
		HasArticles   bool
	}{
		User:          user,
		HasCategories: len(cats) > 0,
		HasArticles:   len(arts) > 0,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

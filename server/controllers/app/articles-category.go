package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type ArticleCategoryController struct {
	Renderer            handler.Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	Category            store.CategoryRepository
	InternalServerError controller.Controller
	NotFound            controller.Controller
}

func (ctr ArticleCategoryController) Serve(ctx context.Context, pg *handler.Page, rqs *http.Request) controller.Controller {
	vars := mux.Vars(rqs)

	info := ctr.SiteInfo.Get(ctx)

	cat, err := ctr.Category.Get(ctx, vars["category"])
	if err != nil {
		log.Errorf(ctx, "app.ArticleCategoryHandler - Unable to get category %s\n%v", vars["category"], err)
		return ctr.NotFound
	}

	arts, err := ctr.Article.GetAllByCategory(ctx, cat, 10)
	if err != nil {
		log.Errorf(ctx, "app.ArticleCategoryHandler - Unable to get articles by category %s\n%v", vars["category"], err)
		return ctr.NotFound
	}

	cnt, err := ctr.Renderer.Render(ctx, "page/app/article", struct {
		Articles []*store.Article
		Category *store.Category
	}{
		Articles: arts,
		Category: cat,
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

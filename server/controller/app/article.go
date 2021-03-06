package app

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

type ArticleController struct {
	Renderer            controller.Renderer
	NotFound            controller.Controller
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	InternalServerError controller.Controller
}

func (ctr ArticleController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	vars := mux.Vars(rqs)

	info := ctr.SiteInfo.Get(ctx)

	cat := store.NewEmptyCategory(vars["category"])

	art, err := ctr.Article.Get(ctx, cat, vars["key"])

	if err != nil {
		log.Errorf(ctx, "app.ArticleHandler - Unable to get articles by type\n%v", err)
		return ctr.NotFound
	}

	cnt, err := ctr.Renderer.String("page/app/article", struct {
		Article *store.Article
	}{
		art,
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

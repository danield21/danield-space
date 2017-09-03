package app

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"golang.org/x/net/context"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
)

type publicationList struct {
	Category *store.Category
	Articles []*store.Article
}

type ArticlesController struct {
	Renderer            Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	InternalServerError controller.Controller
}

func (ctr ArticlesController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	articleMap, err := ctr.Article.GetMapKeyedByCategory(ctx, 3)

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to get articles organized by their type"))
		return ctr.InternalServerError
	}

	var articles []publicationList

	for cat, a := range articleMap {
		articles = append(articles, publicationList{
			Category: cat,
			Articles: a,
		})
	}

	cnt, err := ctr.Renderer.String("page/app/articles", struct {
		Articles []publicationList
	}{
		Articles: articles,
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

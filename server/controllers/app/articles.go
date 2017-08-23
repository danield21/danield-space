package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type publicationList struct {
	Category *store.Category
	Articles []*store.Article
}

type ArticlesHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	Article  store.ArticleRepository
}

func (hnd ArticlesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.New(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	articleMap, err := hnd.Article.GetMapKeyedByCategory(ctx, 10)

	if err != nil {
		log.Errorf(ctx, "app.Articles - Unable to get articles organized by their type\n%v", err)
	}

	var articles []publicationList

	for cat, a := range articleMap {
		articles = append(articles, publicationList{
			Category: cat,
			Articles: a,
		})
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/app/article", struct {
		Articles []publicationList
	}{
		Articles: articles,
	})

	if err != nil {
		log.Errorf(ctx, "app.AboutHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	hnd.Renderer.Send(w, r, pg)
}

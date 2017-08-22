package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type ArticleCategoryHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	Article  store.ArticleRepository
	Category store.CategoryRepository
}

func (hnd ArticleCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx := hnd.Context.New(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Header["description"] = info.ShortDescription()
	pg.Header["author"] = info.Owner

	cat, err := hnd.Category.Get(ctx, vars["category"])
	if err != nil {
		log.Errorf(ctx, "app.ArticlesCategoryPageLink - Unable to get category %s\n%v", vars["category"], err)
		return
	}

	arts, err := hnd.Article.GetAllByCategory(ctx, cat, 1)
	if err != nil {
		log.Errorf(ctx, "app.ArticlesCategoryPageLink - Unable to get articles by category %s\n%v", cat.Title, err)
		return
	}

	pg.Content = template.HTML(hnd.Renderer.Render(ctx, "page/app/articles-type", struct {
		Articles []*store.Article
		Category *store.Category
	}{
		Articles: arts,
		Category: cat,
	}))

	hnd.Renderer.Send(w, r, pg)
}

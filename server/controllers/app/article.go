package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type ArticleHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	NotFound http.Handler
	SiteInfo store.SiteInfoRepository
	Article  store.ArticleRepository
}

func (hnd ArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx := hnd.Context.Generate(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	cat := store.NewEmptyCategory(vars["category"])

	art, err := hnd.Article.Get(ctx, cat, vars["key"])

	if err != nil {
		log.Errorf(ctx, "app.ArticleHandler - Unable to get articles by type\n%v", err)
		hnd.NotFound.ServeHTTP(w, r)
		return
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/app/article", struct {
		Article *store.Article
	}{
		art,
	})

	if err != nil {
		log.Errorf(ctx, "app.ArticleHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	hnd.Renderer.Send(w, r, pg)
}
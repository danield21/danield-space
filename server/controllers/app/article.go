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
	SiteInfo store.SiteInfoRepository
	Article  store.ArticleRepository
}

func (hnd ArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx := hnd.Context.New(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Header["description"] = info.ShortDescription()
	pg.Header["author"] = info.Owner

	cat := store.NewEmptyCategory(vars["category"])

	art, err := hnd.Article.Get(ctx, cat, vars["key"])

	if err != nil {
		log.Errorf(ctx, "app.ArticlePageLink - Unable to get articles by type\n%v", err)
		return
	}

	pg.Content = template.HTML(hnd.Renderer.Render(ctx, "page/app/article", struct {
		Article *store.Article
	}{
		art,
	}))

	hnd.Renderer.Send(w, r, pg)
}

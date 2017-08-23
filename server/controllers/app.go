package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/app"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
)

//App creates a new server instance to run
type AppRouter struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	About    store.AboutRepository
	Article  store.ArticleRepository
	Category store.CategoryRepository
	NotFound http.Handler
}

func (rtr AppRouter) Route(r *mux.Router) {
	r.NotFoundHandler = rtr.NotFound

	r.Handle("/", app.IndexHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
		Article:  rtr.Article,
	})
	r.Handle("/about", app.AboutHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
		About:    rtr.About,
	})
	r.Handle("/articles", app.ArticlesHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
		Article:  rtr.Article,
	})
	r.Handle("/articles/{category}", app.ArticleCategoryHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		NotFound: rtr.NotFound,
		SiteInfo: rtr.SiteInfo,
		Article:  rtr.Article,
		Category: rtr.Category,
	})
	r.Handle("/articles/{category}/{key}", app.ArticleHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		NotFound: rtr.NotFound,
		SiteInfo: rtr.SiteInfo,
		Article:  rtr.Article,
	})
}

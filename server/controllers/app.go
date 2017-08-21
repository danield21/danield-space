package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/app"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func App(e handler.Environment, r *mux.Router) {
	mgr := Migrator{e}
	r.NotFoundHandler = handler.Prepare(e, view.HTMLHandler, handler.ToLink(status.NotFoundPageHandler))

	r.HandleFunc("/", handler.Apply(e, app.IndexHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/", handler.Apply(e, app.IndexPageHandler)).Methods(http.MethodGet)
	r.Handle("/about", app.AboutHandler{
		Context:  mgr,
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
		About:    e.Repository().About(),
	})
	r.HandleFunc("/articles", handler.Apply(e, app.ArticlesHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/articles", handler.Apply(e, app.ArticlesPageHandler)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/articles/{category}", handler.Apply(e, app.ArticlesCategoryHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/articles/{category}", handler.Apply(e, app.ArticlesCategoryPageHandler)).Methods(http.MethodGet, http.MethodPost)
	r.Handle("/articles/{category}/{key}", app.ArticleHandler{
		Context:  mgr,
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
		Article:  e.Repository().Article(),
	})
}

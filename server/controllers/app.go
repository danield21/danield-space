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
	r.NotFoundHandler = handler.Prepare(e, view.HTMLHandler, handler.ToLink(status.NotFoundPageHandler))

	r.HandleFunc("/", handler.Apply(e, app.IndexHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/", handler.Apply(e, app.IndexPageHandler)).Methods(http.MethodGet)
	r.HandleFunc("/about", handler.Apply(e, app.AboutHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/about", handler.Apply(e, app.AboutPageHandler)).Methods(http.MethodGet)
	r.HandleFunc("/articles", handler.Apply(e, app.ArticlesHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/articles", handler.Apply(e, app.ArticlesPageHandler)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/articles/{category}", handler.Apply(e, app.ArticlesCategoryHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/articles/{category}", handler.Apply(e, app.ArticlesCategoryPageHandler)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/articles/{category}/{key}", handler.Apply(e, app.ArticleHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/articles/{category}/{key}", handler.Apply(e, app.ArticlePageHandler)).Methods(http.MethodGet, http.MethodPost)
}

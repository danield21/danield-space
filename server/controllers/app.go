package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/app"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func App(e envir.Environment, r *mux.Router) {
	r.NotFoundHandler = service.Prepare(e, view.HTMLHandler, service.ToLink(status.NotFoundPageHandler), link.Theme)

	r.HandleFunc("/", service.Apply(e, app.IndexHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/", service.Apply(e, app.IndexPageHandler)).Methods(http.MethodGet)
	r.HandleFunc("/publications", service.Apply(e, app.PublicationsHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/publications", service.Apply(e, app.PublicationsPageHandler)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{category}", service.Apply(e, app.PublicationsCategoryHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{category}", service.Apply(e, app.PublicationsCategoryPageHandler)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{category}/{key}", service.Apply(e, app.ArticleHeadersHandler)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{category}/{key}", service.Apply(e, app.ArticlePageHandler)).Methods(http.MethodGet, http.MethodPost)
}

package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/app"
	"github.com/danield21/danield-space/server/handler/status"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func App(e envir.Environment, r *mux.Router) {
	r.HandleFunc("/", handler.Prepare(e, app.IndexHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/", handler.Prepare(e, app.Index)).Methods(http.MethodGet)
	r.HandleFunc("/publications", handler.Prepare(e, app.PublicationsHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications", handler.Prepare(e, app.Publications)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{type}", handler.Prepare(e, app.PublicationsTypeHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}", handler.Prepare(e, app.PublicationsType)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{type}/{key}", handler.Prepare(e, app.ArticleHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}/{key}", handler.Prepare(e, app.Article)).Methods(http.MethodGet, http.MethodPost)
	r.NotFoundHandler = handler.Prepare(e, status.NotFound)
}

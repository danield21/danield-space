package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/app"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/service"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func App(e envir.Environment, r *mux.Router) {
	r.HandleFunc("/", service.Prepare(e, app.IndexHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/", service.Prepare(e, app.Index)).Methods(http.MethodGet)
	r.HandleFunc("/publications", service.Prepare(e, app.PublicationsHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications", service.Prepare(e, app.Publications)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{category}", service.Prepare(e, app.PublicationsTypeHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{category}", service.Prepare(e, app.PublicationsType)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{category}/{key}", service.Prepare(e, app.ArticleHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{category}/{key}", service.Prepare(e, app.Article)).Methods(http.MethodGet, http.MethodPost)
}

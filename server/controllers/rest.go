package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/rest/account"
	"github.com/danield21/danield-space/server/controllers/rest/article"
	"github.com/danield21/danield-space/server/controllers/rest/category"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"github.com/gorilla/mux"
)

//Rest configures the services for rest services
func Rest(e envir.Environment, r *mux.Router) {
	r.NotFoundHandler = service.Prepare(e, view.JSONHandler, status.NotFoundLink)

	r.HandleFunc("/article", service.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/article/{category}/{url}", service.Prepare(e, article.Put)).Methods(http.MethodPut)
	r.HandleFunc("/category", service.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/category/{category}", service.Prepare(e, category.Put)).Methods(http.MethodPut)
	r.HandleFunc("/admin/authenticate", service.Prepare(e, account.Auth)).Methods(http.MethodPost)
	r.HandleFunc("/admin/unauthenticate", service.Prepare(e, account.Unauth)).Methods(http.MethodPost)
}

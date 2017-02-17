package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/rest/account"
	"github.com/danield21/danield-space/server/handler/rest/article"
	"github.com/danield21/danield-space/server/handler/rest/category"
	"github.com/gorilla/mux"
)

//Rest configures the handlers for rest services
func Rest(e envir.Environment, r *mux.Router) {
	r.HandleFunc("/article", handler.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/article/{category}/{url}", handler.Prepare(e, article.Put)).Methods(http.MethodPut)
	r.HandleFunc("/category", handler.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/category/{category}", handler.Prepare(e, category.Put)).Methods(http.MethodPut)
	r.HandleFunc("/admin/authenticate", handler.Prepare(e, account.Auth)).Methods(http.MethodPost)
	r.HandleFunc("/admin/unauthenticate", handler.Prepare(e, account.Unauth)).Methods(http.MethodPost)
}

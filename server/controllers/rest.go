package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/rest/account"
	"github.com/danield21/danield-space/server/controllers/rest/article"
	"github.com/danield21/danield-space/server/controllers/rest/category"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/gorilla/mux"
)

//Rest configures the handlers for rest handlers
func Rest(e handler.Environment, r *mux.Router) {
	r.NotFoundHandler = handler.Prepare(e, view.JSONHandler, handler.ToLink(status.NotFoundPageHandler))

	r.HandleFunc("/article", handler.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/article/{category}/{url}", handler.Prepare(e, article.Put)).Methods(http.MethodPut)
	r.HandleFunc("/category", handler.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/category/{category}", handler.Prepare(e, category.Put)).Methods(http.MethodPut)
	r.HandleFunc("/admin/authenticate", handler.Prepare(e, account.Auth)).Methods(http.MethodPost)
	r.HandleFunc("/admin/unauthenticate", handler.Prepare(e, account.Unauth)).Methods(http.MethodPost)
}

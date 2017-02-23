package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/gorilla/mux"
)

//Admin configures the handlers for admin handlers
func Admin(e handler.Environment, r *mux.Router) {
	r.NotFoundHandler = handler.Prepare(e, view.HTMLHandler, handler.ToLink(status.NotFoundPageHandler), link.Theme)

	r.HandleFunc("/", handler.Apply(e, admin.IndexHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/", handler.Apply(e, admin.IndexPageHandler)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/signin", handler.Apply(e, admin.SignInHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/signin", handler.Apply(e, admin.SignInPageHandler)).
		Methods(http.MethodGet)
	r.HandleFunc("/signin", handler.Apply(e, admin.SignInActionHandler)).
		Methods(http.MethodPost)
	r.HandleFunc("/sign-out", handler.Apply(e, admin.SignOutHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/sign-out", handler.Apply(e, admin.SignOutPageHandler)).
		Methods(http.MethodGet)
	r.HandleFunc("/sign-out", handler.Apply(e, admin.SignOutActionHandler)).
		Methods(http.MethodPost)
	r.HandleFunc("/article/", handler.Apply(e, admin.ArticlesHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/", handler.Apply(e, admin.ArticlesPageHandler)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/publish", handler.Apply(e, admin.PublishHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/publish", handler.Apply(e, admin.PublishPageHandler)).
		Methods(http.MethodGet)
	r.HandleFunc("/article/publish", handler.Apply(e, admin.PublishActionHandler)).
		Methods(http.MethodPost)
	r.HandleFunc("/account/", handler.Apply(e, admin.AccountHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/account/", handler.Apply(e, admin.AccountPageHandler)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/site-info/", handler.Apply(e, admin.SiteInfoHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/site-info/", handler.Apply(e, admin.SiteInfoPageHandler)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/", handler.Apply(e, admin.CategoryHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/", handler.Apply(e, admin.CategoryPageHandler)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/create", handler.Apply(e, admin.CategoryCreateHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/create", handler.Apply(e, admin.CategoryCreatePageHandler)).
		Methods(http.MethodGet)
	r.HandleFunc("/category/create", handler.Apply(e, admin.CategoryCreateActionHandler)).
		Methods(http.MethodPost)
}

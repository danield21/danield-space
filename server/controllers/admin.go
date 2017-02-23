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
	r.HandleFunc("/article/", handler.Prepare(e, admin.ArticlesHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/", handler.Prepare(e, admin.ArticlesPageHandler, status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.PublishHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.PublishPageHandler, status.LinkAll)).
		Methods(http.MethodGet)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.PublishFormHandler, status.LinkAll)).
		Methods(http.MethodPost)
	r.HandleFunc("/account/", handler.Prepare(e, admin.AccountHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/account/", handler.Prepare(e, admin.AccountPageHandler, status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/site-info/", handler.Prepare(e, admin.SiteInfoHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/site-info/", handler.Prepare(e, admin.SiteInfoPageHandler, status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/", handler.Prepare(e, admin.CategoryHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/", handler.Prepare(e, admin.CategoryPageHandler, status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/create", handler.Apply(e, admin.CategoryCreateHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/create", handler.Apply(e, admin.CategoryCreatePageHandler)).
		Methods(http.MethodGet)
	r.HandleFunc("/category/create", handler.Apply(e, admin.CategoryCreateActionHandler)).
		Methods(http.MethodPost)
}

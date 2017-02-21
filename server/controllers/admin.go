package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"github.com/gorilla/mux"
)

//Admin configures the services for admin services
func Admin(e envir.Environment, r *mux.Router) {
	r.NotFoundHandler = service.Prepare(e, view.HTMLHandler, service.ToLink(status.NotFoundPageHandler), link.Theme)

	r.HandleFunc("/", service.Apply(e, admin.IndexHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/", service.Apply(e, admin.IndexPageHandler)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/signin", service.Prepare(e, admin.SignInHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/signin", service.Prepare(e, admin.SignIn)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/", service.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/", service.Prepare(e, admin.ShowPage("article"), status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/publish", service.Prepare(e, admin.PublishHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/publish", service.Prepare(e, admin.PublishPageHandler, status.LinkAll)).
		Methods(http.MethodGet)
	r.HandleFunc("/article/publish", service.Prepare(e, admin.PublishFormHandler, status.LinkAll)).
		Methods(http.MethodPost)
	r.HandleFunc("/account/", service.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/account/", service.Prepare(e, admin.ShowPage("account"), status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/site-info/", service.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/site-info/", service.Prepare(e, admin.ShowPage("site-info"), status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/", service.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/", service.Prepare(e, admin.ShowPage("category"), status.LinkAll)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/create", service.Apply(e, admin.CategoryCreateHeadersHandler)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/create", service.Apply(e, admin.CategoryCreatePageHandler)).
		Methods(http.MethodGet)
	r.HandleFunc("/category/create", service.Apply(e, admin.CategoryCreateFormHandler)).
		Methods(http.MethodPost)
}

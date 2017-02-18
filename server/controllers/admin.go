package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/admin"
	"github.com/gorilla/mux"
)

//Admin configures the handlers for admin services
func Admin(e envir.Environment, r *mux.Router) {
	r.HandleFunc("/", handler.Prepare(e, admin.IndexHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/", handler.Prepare(e, admin.Index, admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/signin", handler.Prepare(e, admin.SignInHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/signin", handler.Prepare(e, admin.SignIn)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/", handler.Prepare(e, admin.ShowPage("article"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.PublishHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.Publish, admin.Authorized)).
		Methods(http.MethodGet)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.ArticlePublish, admin.Authorized)).
		Methods(http.MethodPost)
	r.HandleFunc("/account/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/account/", handler.Prepare(e, admin.ShowPage("account"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/site-info/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/site-info/", handler.Prepare(e, admin.ShowPage("site-info"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/", handler.Prepare(e, admin.ShowPage("category"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/create", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/create", handler.Prepare(e, admin.ShowPage("category-create"), admin.Authorized)).
		Methods(http.MethodGet)
	r.HandleFunc("/category/create", handler.Prepare(e, admin.CategoryCreate, admin.Authorized)).
		Methods(http.MethodPost)
}

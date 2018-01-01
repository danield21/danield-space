package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/router"
	"github.com/danield21/danield-space/server/views"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func New() http.Handler {
	r := mux.NewRouter()

	connections := CreateRepository()
	session := SessionGenerator{
		connections.Session,
	}
	context := ContextGenerater{}

	rnd := views.Get("view")

	router.AppRouter{
		Context:  context,
		Session:  session,
		Renderer: rnd,
		SiteInfo: connections.SiteInfo,
		About:    connections.About,
		Article:  connections.Article,
		Category: connections.Category,
	}.Route(r)

	router.AdminRouter{
		Context:  context,
		Session:  session,
		Renderer: rnd,
		SiteInfo: connections.SiteInfo,
		About:    connections.About,
		Article:  connections.Article,
		Category: connections.Category,
	}.Route(r.PathPrefix("/admin").Subrouter())

	return r
}

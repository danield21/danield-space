package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers"
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

	controllers.AppRouter{
		Context:  context,
		Session:  session,
		Renderer: rnd,
		SiteInfo: connections.SiteInfo,
		About:    connections.About,
		Article:  connections.Article,
		Category: connections.Category,
	}.Route(r)

	controllers.AdminRouter{
		Context:  context,
		Session:  session,
		Renderer: rnd,
		SiteInfo: connections.SiteInfo,
		Account:  connections.Account,
		About:    connections.About,
		Article:  connections.Article,
		Category: connections.Category,
	}.Route(r.PathPrefix("/admin").Subrouter())

	return r
}

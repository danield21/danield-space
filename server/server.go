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
	e := ProductionEnvironment{
		Connections: CreateRepository(),
	}

	rnd := NewRenderer(views.Get("view"))

	mgr := controllers.Migrator{
		Environment: &e,
	}

	controllers.AppRouter{
		Context:  mgr.Context(),
		Session:  mgr.Session(),
		Renderer: rnd,
		SiteInfo: e.Repository().SiteInfo(),
		About:    e.Repository().About(),
		Article:  e.Repository().Article(),
		Category: e.Repository().Category(),
	}.Route(r)

	controllers.AdminRouter{
		Context:  mgr.Context(),
		Session:  mgr.Session(),
		Renderer: rnd,
		SiteInfo: e.Repository().SiteInfo(),
		Account:  e.Repository().Account(),
		About:    e.Repository().About(),
		Article:  e.Repository().Article(),
		Category: e.Repository().Category(),
	}.Route(r.PathPrefix("/admin").Subrouter())

	return r
}

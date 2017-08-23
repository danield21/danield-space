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
		GenerateTemplates: views.Get("view"),
		Connections:       CreateRepository(),
	}

	mgr := controllers.Migrator{
		Environment: &e,
	}

	controllers.AppRouter{
		Context:  mgr,
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
		About:    e.Repository().About(),
		Article:  e.Repository().Article(),
		Category: e.Repository().Category(),
	}.Route(r)
	controllers.Admin(&e, r.PathPrefix("/admin").Subrouter())

	return r
}

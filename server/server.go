package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers"
	"github.com/danield21/danield-space/server/controllers/status"
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

	notFnd := status.NotFoundHandler{
		Context:  mgr.Context(),
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
	}

	controllers.AppRouter{
		Context:  mgr.Context(),
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
		About:    e.Repository().About(),
		Article:  e.Repository().Article(),
		Category: e.Repository().Category(),
		NotFound: notFnd,
	}.Route(r)
	controllers.Admin(&e, r.PathPrefix("/admin").Subrouter())

	return r
}

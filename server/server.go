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

	unauth := status.UnauthorizedHandler{
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

	controllers.AdminRouter{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Account:      e.Repository().Account(),
		About:        e.Repository().About(),
		Article:      e.Repository().Article(),
		Category:     e.Repository().Category(),
		NotFound:     notFnd,
		Unauthorized: unauth,
	}.Route(r.PathPrefix("/admin").Subrouter())

	return r
}

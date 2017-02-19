package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/status"
	"github.com/danield21/danield-space/server/views"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func New() http.Handler {
	r := mux.NewRouter()
	e := ProductionEnvironment{Templates: views.Get()}

	controllers.App(e, r)
	controllers.Admin(e, r.PathPrefix("/admin").Subrouter())
	controllers.Rest(e, r.PathPrefix("/rest").Subrouter())
	r.NotFoundHandler = service.Apply(e, status.NotFoundHandler)

	return r
}

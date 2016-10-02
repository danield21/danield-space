package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/handlers"
	"github.com/danield21/danield-space/server/views"
	"github.com/gorilla/mux"
)

//New creates a new server instance to run
func New() http.Handler {
	r := mux.NewRouter()
	config := handlers.Config{Views: views.Get()}

	r.HandleFunc("/", handlers.Index(config)).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.IndexHeaders(config)).Methods(http.MethodOptions)

	return r
}

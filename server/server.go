package server

import (
	"github.com/danield21/danield-space/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/danield21/danield-space/server/views"
)

func New() http.Handler {
	r := mux.NewRouter()
	config := handlers.Config{Views: views.Get()}

	r.HandleFunc("/", handlers.Index(config)).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.IndexHeaders(config)).Methods(http.MethodOptions)

	return r
}
package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/danield21/danield-space/server/handlers"
	"github.com/danield21/danield-space/server/views"
	"github.com/gorilla/mux"
)

type config struct {
	Templates *template.Template
}

func (c config) View(w io.Writer, view string, data interface{}) error {
	return c.Templates.ExecuteTemplate(w, view, data)
}

//New creates a new server instance to run
func New() http.Handler {
	r := mux.NewRouter()
	c := config{Templates: views.Get()}

	r.HandleFunc("/", handlers.IndexHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/", handlers.Index(c)).Methods(http.MethodGet)
	r.HandleFunc("/publications/{type}", handlers.PublicationsTypeHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}", handlers.PublicationsType(c)).Methods(http.MethodGet)
	r.HandleFunc("/publications/{type}/{key}", handlers.ArticleHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}/{key}", handlers.Article(c)).Methods(http.MethodGet)
	r.NotFoundHandler = handlers.NotFound(c)

	return r
}

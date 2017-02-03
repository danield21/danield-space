package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/handlers"
	"github.com/danield21/danield-space/server/handlers/admin"
	rAdmin "github.com/danield21/danield-space/server/handlers/rest/admin"
	"github.com/danield21/danield-space/server/handlers/rest/article"
	"github.com/danield21/danield-space/server/views"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type settings struct {
	Templates *template.Template
}

func (c settings) View(w io.Writer, theme, view string, data interface{}) error {
	return config.RenderTemplateWithTheme(c.Templates, w, theme, view, data)
}

func (c settings) GetSession(r *http.Request) (session *sessions.Session) {
	return config.GetSession(r)
}

//New creates a new server instance to run
func New() http.Handler {
	r := mux.NewRouter()
	c := settings{Templates: views.Get()}

	r.HandleFunc("/", handlers.IndexHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/", handlers.Index(c)).Methods(http.MethodGet)
	r.HandleFunc("/publications", handlers.PublicationsHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/publications", handlers.Publications(c)).Methods(http.MethodGet)
	r.HandleFunc("/publications/{type}", handlers.PublicationsTypeHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}", handlers.PublicationsType(c)).Methods(http.MethodGet)
	r.HandleFunc("/publications/{type}/{key}", handlers.ArticleHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}/{key}", handlers.Article(c)).Methods(http.MethodGet)
	r.NotFoundHandler = handlers.NotFound(c)

	Admin(r.PathPrefix("/admin").Subrouter(), c)
	Rest(r.PathPrefix("/rest").Subrouter(), c)

	return r
}

func Rest(r *mux.Router, c config.Config) {
	r.HandleFunc("/article", article.Get(c)).Methods(http.MethodGet)
	r.HandleFunc("/admin/authenticate", rAdmin.Auth(c)).Methods(http.MethodPost)
}

func Admin(r *mux.Router, c config.Config) {
	r.HandleFunc("/publish", admin.PublishHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/publish", admin.Publish(c)).Methods(http.MethodGet)
	r.HandleFunc("/signin", admin.SignInHeaders(c)).Methods(http.MethodHead)
	r.HandleFunc("/signin", admin.SignIn(c)).Methods(http.MethodGet)
}

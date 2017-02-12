package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"github.com/danield21/danield-space/pkg/handler/admin"
	"github.com/danield21/danield-space/pkg/handler/app"
	"github.com/danield21/danield-space/pkg/handler/rest/account"
	"github.com/danield21/danield-space/pkg/handler/rest/article"
	"github.com/danield21/danield-space/pkg/handler/status"
	"github.com/danield21/danield-space/pkg/views"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//ProductionEnvironment contains all the data required to run the server
type ProductionEnvironment struct {
	Templates *template.Template
}

//View generates a view based on the templates stored
func (p ProductionEnvironment) View(w io.Writer, theme, view string, data interface{}) error {
	return envir.RenderTemplateWithTheme(p.Templates, w, theme, view, data)
}

//Session gets the session using a secure key
func (p ProductionEnvironment) Session(r *http.Request) (session *sessions.Session) {
	session = envir.GetSession(r)
	return
}

//Theme gets the requested theme
func (p ProductionEnvironment) Theme(r *http.Request, defaultTheme string) (theme string) {
	theme = envir.GetTheme(r, defaultTheme)
	return
}

//Context gets a new app engine context
func (p ProductionEnvironment) Context(r *http.Request) (ctx context.Context) {
	ctx = appengine.NewContext(r)
	return
}

//New creates a new server instance to run
func New() http.Handler {
	r := mux.NewRouter()
	e := ProductionEnvironment{Templates: views.Get()}

	r.HandleFunc("/", handler.Prepare(app.IndexHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/", handler.Prepare(app.Index, e)).Methods(http.MethodGet)
	r.HandleFunc("/publications", handler.Prepare(app.PublicationsHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/publications", handler.Prepare(app.Publications, e)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{type}", handler.Prepare(app.PublicationsTypeHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}", handler.Prepare(app.PublicationsType, e)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{type}/{key}", handler.Prepare(app.ArticleHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}/{key}", handler.Prepare(app.Article, e)).Methods(http.MethodGet, http.MethodPost)
	r.NotFoundHandler = handler.Prepare(status.NotFound, e)

	Admin(r.PathPrefix("/admin").Subrouter(), e)
	Rest(r.PathPrefix("/rest").Subrouter(), e)

	return r
}

//Rest configures the handlers for rest services
func Rest(r *mux.Router, e ProductionEnvironment) {
	r.HandleFunc("/article", handler.Prepare(article.Get, e)).Methods(http.MethodGet)
	r.HandleFunc("/admin/authenticate", handler.Prepare(account.Auth, e)).Methods(http.MethodPost)
	r.HandleFunc("/admin/unauthenticate", handler.Prepare(account.Unauth, e)).Methods(http.MethodPost)
}

//Admin configures the handlers for admin services
func Admin(r *mux.Router, e ProductionEnvironment) {
	r.HandleFunc("/", handler.Prepare(admin.IndexHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/", handler.Prepare(admin.Index, e)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/publish", handler.Prepare(admin.PublishHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/article/publish", handler.Prepare(admin.Publish, e)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/signin", handler.Prepare(admin.SignInHeaders, e)).Methods(http.MethodHead)
	r.HandleFunc("/signin", handler.Prepare(admin.SignIn, e)).Methods(http.MethodGet, http.MethodPost)
}

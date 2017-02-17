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
	"github.com/danield21/danield-space/pkg/handler/rest/category"
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

	r.HandleFunc("/", handler.Prepare(e, app.IndexHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/", handler.Prepare(e, app.Index)).Methods(http.MethodGet)
	r.HandleFunc("/publications", handler.Prepare(e, app.PublicationsHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications", handler.Prepare(e, app.Publications)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{type}", handler.Prepare(e, app.PublicationsTypeHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}", handler.Prepare(e, app.PublicationsType)).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/publications/{type}/{key}", handler.Prepare(e, app.ArticleHeaders)).Methods(http.MethodHead)
	r.HandleFunc("/publications/{type}/{key}", handler.Prepare(e, app.Article)).Methods(http.MethodGet, http.MethodPost)
	r.NotFoundHandler = handler.Prepare(e, status.NotFound)

	Admin(r.PathPrefix("/admin").Subrouter(), e)
	Rest(r.PathPrefix("/rest").Subrouter(), e)

	return r
}

//Rest configures the handlers for rest services
func Rest(r *mux.Router, e ProductionEnvironment) {
	r.HandleFunc("/article", handler.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/article/{category}/{url}", handler.Prepare(e, article.Put)).Methods(http.MethodPut)
	r.HandleFunc("/category", handler.Prepare(e, article.Get)).Methods(http.MethodGet)
	r.HandleFunc("/category/{category}", handler.Prepare(e, category.Put)).Methods(http.MethodPut)
	r.HandleFunc("/admin/authenticate", handler.Prepare(e, account.Auth)).Methods(http.MethodPost)
	r.HandleFunc("/admin/unauthenticate", handler.Prepare(e, account.Unauth)).Methods(http.MethodPost)
}

//Admin configures the handlers for admin services
func Admin(r *mux.Router, e ProductionEnvironment) {
	r.HandleFunc("/", handler.Prepare(e, admin.IndexHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/", handler.Prepare(e, admin.Index, admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/signin", handler.Prepare(e, admin.SignInHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/signin", handler.Prepare(e, admin.SignIn)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/", handler.Prepare(e, admin.ShowPage("article"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.PublishHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/article/publish", handler.Prepare(e, admin.Publish, admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/account/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/account/", handler.Prepare(e, admin.ShowPage("account"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/site-info/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/site-info/", handler.Prepare(e, admin.ShowPage("site-info"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/", handler.Prepare(e, admin.ShowPage("category"), admin.Authorized)).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/category/create", handler.Prepare(e, admin.ShowPageHeaders)).
		Methods(http.MethodHead)
	r.HandleFunc("/category/create", handler.Prepare(e, admin.ShowPage("category-create"), admin.Authorized)).
		Methods(http.MethodGet)
	r.HandleFunc("/category/create", handler.Prepare(e, admin.CategoryCreate, admin.Authorized)).
		Methods(http.MethodPost)
}

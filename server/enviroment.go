package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
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
	return handler.RenderTemplateWithTheme(p.Templates, w, theme, view, data)
}

//Session gets the session using a secure key
func (p ProductionEnvironment) Session(r *http.Request) (session *sessions.Session) {
	session = handler.GetSession(r)
	return
}

//Context gets a new app engine context
func (p ProductionEnvironment) Context(r *http.Request) (ctx context.Context) {
	ctx = appengine.NewContext(r)
	return
}

package server

import (
	"html/template"
	"io"
	"net/http"

	"google.golang.org/appengine"

	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

//ProductionEnvironment contains all the data required to run the server
type ProductionEnvironment struct {
	Templates         *template.Template
	GenerateTemplates <-chan *template.Template
}

//View generates a view based on the templates stored
func (p *ProductionEnvironment) View(w io.Writer, theme, view string, data interface{}) error {
	if p.Templates == nil {
		p.Templates = <-p.GenerateTemplates
	}

	return RenderTemplateWithTheme(p.Templates, w, theme, view, data)
}

//Session gets the session using a secure key
func (p *ProductionEnvironment) Session(r *http.Request) (session *sessions.Session) {
	session = GetSession(r)
	return
}

//Context gets a new app engine context
func (p *ProductionEnvironment) Context(r *http.Request) (ctx context.Context) {
	ctx = appengine.NewContext(r)
	return
}

package controllers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

//TestingEnvironment is only used for unit tests
type TestingEnvironment struct {
	Templates *template.Template
	Ctx       context.Context
}

//View creates a mock view
func (t TestingEnvironment) View(w io.Writer, theme, view string, data interface{}) error {
	return handler.RenderTemplateWithTheme(t.Templates, w, theme, view, data)
}

//Session gets a mock session
func (t TestingEnvironment) Session(r *http.Request) (session *sessions.Session) {
	session = handler.GetSession(r)
	return
}

//Context gets a mock context
func (t TestingEnvironment) Context(r *http.Request) (ctx context.Context) {
	ctx = t.Ctx
	return
}

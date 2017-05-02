package server

import (
	"html/template"
	"io"
	"net/http"

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
	return RenderTemplateWithTheme(t.Templates, w, theme, view, data)
}

//Session gets a mock session
func (t TestingEnvironment) Session(r *http.Request) *sessions.Session {
	return GetSession(r)
}

//Context gets a mock context
func (t TestingEnvironment) Context(r *http.Request) context.Context {
	return t.Ctx
}

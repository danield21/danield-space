package server

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
func (t TestingEnvironment) View(w io.Writer, view string, data interface{}) error {
	return RenderTemplate(t.Templates, w, view, data)
}

//View creates a mock view
func (t TestingEnvironment) Partial(w io.Writer, view string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, view, data)
}

//Session gets a mock session
func (t TestingEnvironment) Session(ctx context.Context, r *http.Request) *sessions.Session {
	return GetSession(ctx, t, r)
}

//Context gets a mock context
func (t TestingEnvironment) Context(r *http.Request) context.Context {
	return t.Ctx
}

func (t TestingEnvironment) Repository() handler.Repositories {
	return nil
}

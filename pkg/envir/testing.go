package envir

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
func (t TestingEnvironment) Session(r *http.Request) (session *sessions.Session) {
	session = GetSession(r)
	return
}

//Theme gets a mock theme
func (t TestingEnvironment) Theme(r *http.Request) (theme string) {
	theme = GetTheme(r)
	return
}

//Context gets a mock context
func (t TestingEnvironment) Context(r *http.Request) (ctx context.Context) {
	ctx = t.Ctx
	return
}

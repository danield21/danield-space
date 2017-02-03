package config

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
)

//MockConfig is only used for unit tests
type MockConfig struct {
	Templates *template.Template
}

//View creates a mock view
func (c MockConfig) View(w io.Writer, theme, view string, data interface{}) error {
	return RenderTemplateWithTheme(c.Templates, w, theme, view, data)
}

//Gets a mock session
func (c MockConfig) GetSession(r *http.Request) (session *sessions.Session) {
	return GetSession(r)
}

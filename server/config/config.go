package config

import (
	"io"
	"net/http"

	"github.com/gorilla/sessions"
)

//Config holds views and database connections to inject into handlers
type Config interface {
	View(w io.Writer, theme, view string, data interface{}) error
	GetSession(r *http.Request) (session *sessions.Session)
}

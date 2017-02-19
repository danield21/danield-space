package envir

import (
	"io"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

//Environment holds all information that a service may need
type Environment interface {
	View(w io.Writer, theme, view string, data interface{}) error
	Context(r *http.Request) context.Context
	//Removing soon
	Session(r *http.Request) (session *sessions.Session)
	Theme(r *http.Request, defaultTheme string) string
}

type Scope interface {
	Session() *sessions.Session
	Request() *http.Request
}

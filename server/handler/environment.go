package handler

import (
	"io"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

//Environment holds all information that a handler may need
type Environment interface {
	View(w io.Writer, theme, view string, data interface{}) error
	Context(r *http.Request) context.Context
	Session(r *http.Request) *sessions.Session
}

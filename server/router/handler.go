package router

import (
	"context"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionGenerator interface {
	Generate(ctx context.Context, rqs *http.Request) *sessions.Session
}

type ContextGenerator interface {
	Generate(rqs *http.Request) context.Context
}

type Renderer interface {
	Render(w io.Writer, view string, data interface{}) error
	String(view string, data interface{}) (string, error)
}

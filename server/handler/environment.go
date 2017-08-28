package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

//Environment holds all information that a handler may need
type Environment interface {
	Context(r *http.Request) context.Context
	Session(ctx context.Context, r *http.Request) *sessions.Session
	Repository() Repositories
}

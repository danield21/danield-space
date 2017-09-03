package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

type SessionGenerator interface {
	Generate(ctx context.Context, rqs *http.Request) *sessions.Session
}

type SessionGeneratorFunc func(ctx context.Context, rqs *http.Request) *sessions.Session

func (sesGen SessionGeneratorFunc) Generate(ctx context.Context, rqs *http.Request) *sessions.Session {
	return sesGen(ctx, rqs)
}

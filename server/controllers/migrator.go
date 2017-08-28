package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

type Migrator struct {
	handler.Environment
}

func (mgr Migrator) Context() handler.ContextGenerator {
	return handler.ContextGeneratorFunc(func(r *http.Request) context.Context {
		return mgr.Environment.Context(r)
	})
}

func (mgr Migrator) Session() handler.SessionGenerator {
	return handler.SessionGeneratorFunc(func(ctx context.Context, r *http.Request) *sessions.Session {
		return mgr.Environment.Session(ctx, r)
	})
}

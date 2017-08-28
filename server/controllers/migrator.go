package controllers

import (
	"bytes"
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

func (mgr Migrator) Render(ctx context.Context, view string, data interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := mgr.Environment.Partial(buffer, view, data)
	return buffer.Bytes(), err
}

func (mgr Migrator) Send(w http.ResponseWriter, r *http.Request, pg *handler.Page) error {
	if pg.Status != 0 {
		w.WriteHeader(pg.Status)
	}

	for header, value := range pg.Header {
		w.Header().Add(header, value)
	}

	return mgr.Environment.Partial(w, "core/page", pg)
}

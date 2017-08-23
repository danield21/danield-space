package controllers

import (
	"bytes"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
)

type Migrator struct {
	handler.Environment
}

func (mgr Migrator) New(r *http.Request) context.Context {
	return mgr.Environment.Context(r)
}

func (mgr Migrator) Render(ctx context.Context, view string, data interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := mgr.Environment.Partial(buffer, view, data)
	return buffer.Bytes(), err
}

func (mgr Migrator) Send(w http.ResponseWriter, r *http.Request, pg *handler.Page) error {
	return mgr.Environment.Partial(w, "core/page", pg)
}

package handler

import (
	"net/http"

	"golang.org/x/net/context"
)

type Renderer interface {
	Render(ctx context.Context, view string, data interface{}) ([]byte, error)
	Send(w http.ResponseWriter, r *http.Request, pg *Page) error
}

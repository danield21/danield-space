package view

import (
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
)

const HTMLContentType = "text/html; charset=utf-8"
const JSONContentType = "application/json; charset=utf-8"

type Header struct {
	Header string
	Value  string
}

func HeaderHandler(status int, headers ...Header) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		w.WriteHeader(status)
		for _, h := range headers {
			w.Header().Add(h.Header, h.Value)
		}
		return ctx, nil
	}
}

package view

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/service"
	"golang.org/x/net/context"
)

type Header struct {
	Header string
	Value  string
}

func HeaderHandler(status int, headers ...Header) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		w.WriteHeader(status)
		for _, h := range headers {
			w.Header().Add(h.Header, h.Value)
		}
		return ctx, nil
	}
}

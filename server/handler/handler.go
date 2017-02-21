package handler

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//Handler is a modified struct of http.HandlerFunc, except requires a Environment for getting information about the site.
type Handler func(ctx context.Context, e Environment, w http.ResponseWriter) (context.Context, error)

type Link func(h Handler) Handler

func Prepare(e Environment, h Handler, links ...Link) http.HandlerFunc {
	return Apply(e, Chain(h, links...))
}

//Apply will turn a Handler into a HandlerFunc with an Environment, may be production or testing
func Apply(e Environment, h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := e.Context(r)
		ctx = SetupContext(ctx, e, r)

		_, err := h(ctx, e, w)

		if err == nil {
			return
		}

		log.Errorf(ctx, "Unexpected error while handling %s\n%v", r.URL.String(), err)
	}
}

func Chain(h Handler, links ...Link) (chain Handler) {
	chain = h
	for _, link := range links {
		chain = link(chain)
	}
	return
}

func ToLink(l Handler) Link {
	return func(h Handler) Handler {
		return func(ctx context.Context, e Environment, w http.ResponseWriter) (context.Context, error) {
			var err error
			ctx, err = l(ctx, e, w)
			if err != nil {
				return ctx, err
			}
			return h(ctx, e, w)
		}
	}
}

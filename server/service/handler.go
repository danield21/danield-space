package service

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Handler is a modified struct of http.HandlerFunc, except requires a Environment for getting information about the site.
type Handler func(s envir.Scope, e envir.Environment, w http.ResponseWriter) error

type Link func(h Handler) Handler

func Prepare(e envir.Environment, h Handler, links ...Link) http.HandlerFunc {
	return Apply(e, Chain(h, links...))
}

//Apply will turn a Handler into a HandlerFunc with an Environment, may be production or testing
func Apply(e envir.Environment, h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scp := NewScope(r)

		err := h(scp, e, w)

		if err == nil {
			return
		}

		ctx := appengine.NewContext(r)

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
		return func(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
			err := l(scp, e, w)
			if err != nil {
				return err
			}
			return h(scp, e, w)
		}
	}
}

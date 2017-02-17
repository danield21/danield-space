package handler

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
)

//Handler is a modified struct of http.HandlerFunc, except requires a Environment for getting information about the site.
type Handler func(e envir.Environment, w http.ResponseWriter, r *http.Request)

type Link func(h Handler) Handler

func Prepare(e envir.Environment, h Handler, links ...Link) http.HandlerFunc {
	return Apply(e, Chain(h, links...))
}

//Apply will turn a Handler into a HandlerFunc with an Environment, may be production or testing
func Apply(e envir.Environment, h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(e, w, r)
	}
}

func Chain(h Handler, links ...Link) (chain Handler) {
	chain = h
	for _, link := range links {
		chain = link(chain)
	}
	return
}

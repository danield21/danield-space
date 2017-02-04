package handler

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/envir"
)

//Handler is a modified struct of http.HandlerFunc, except requires a Environment for getting information about the site.
type Handler func(e envir.Environment, w http.ResponseWriter, r *http.Request)

//Prepare will turn a Handler into a HandlerFunc with an Environment, may be production or testing
func Prepare(h Handler, e envir.Environment) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(e, w, r)
	}
}

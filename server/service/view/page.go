package view

import (
	"errors"
	"net/http"

	"github.com/danield21/danield-space/server/envir"
)

type HTML interface {
	Theme() string
	Page() string
	Data() interface{}
}

var ErrIncompatibleType = errors.New("was not passed value that implemented HTML inteface")

func HTMLHandler(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
	html, ok := scp.(HTML)
	if !ok {
		return ErrIncompatibleType
	}

	return e.View(w, html.Theme(), html.Page(), html.Data())
}

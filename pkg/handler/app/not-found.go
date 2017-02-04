package app

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
)

//NotFound handles the not found page
func NotFound(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusNotFound)

	theme := e.Theme(r)

	err := e.View(w, theme, "page/not-found", nil)
	if err != nil {
		log.Print(err)
	}
}

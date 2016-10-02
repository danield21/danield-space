package handlers

import (
	"github.com/danield21/danield-space/server/content"
	"net/http"
)

func IndexHeaders(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

func Index(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		IndexHeaders(c)(w, r)
		c.Views.ExecuteTemplate(w, "pages/index", nil)
	}
}

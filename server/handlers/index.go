package handlers

import (
	"net/http"
	"github.com/danield21/danield-space/server/content"
)

func IndexHeaders(c Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.Html.AddCharset("utf-8").String())
	}
}

func Index(c Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		IndexHeaders(c)(w, r)
		c.Views.ExecuteTemplate(w, "pages/index", nil)
	}
}

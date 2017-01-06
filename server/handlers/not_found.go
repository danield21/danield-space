package handlers

import (
	"net/http"

	"log"

	"github.com/danield21/danield-space/server/content"
)

//NotFound handles the not found page
func NotFound(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
		w.WriteHeader(http.StatusNotFound)

		err := c.View(w, "pages/not_found", nil)
		if err != nil {
			log.Print(err)
		}
	}
}

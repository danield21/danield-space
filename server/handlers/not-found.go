package handlers

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
)

//NotFound handles the not found page
func NotFound(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
		w.WriteHeader(http.StatusNotFound)

		theme := config.GetTheme(r)

		err := c.View(w, theme, "page/not-found", nil)
		if err != nil {
			log.Print(err)
		}
	}
}

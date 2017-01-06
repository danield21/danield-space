package handlers

import (
	"net/http"

	"log"

	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers"
)

type indexModel struct {
	SiteInfo controllers.SiteInfo
}

//IndexHeaders contains the headers for index
func IndexHeaders(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//Index handles the index page
func Index(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		siteInfo := controllers.SiteInfoController{}

		pageData := indexModel{
			SiteInfo: siteInfo.Get(),
		}

		IndexHeaders(c)(w, r)
		err := c.View(w, "pages/index", pageData)
		if err != nil {
			log.Print(err)
		}
	}
}

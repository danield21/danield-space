package handlers

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers"
	"github.com/gorilla/mux"
)

type publicationsTypeModel struct {
	SiteInfo controllers.SiteInfo
	Articles []controllers.Article
	Type     string
}

//PublicationsTypeHeaders contains the headers for index
func PublicationsTypeHeaders(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//PublicationsType handles the index page
func PublicationsType(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		siteInfo := controllers.SiteInfoController{}
		articles := controllers.ArticleController{}

		pageData := publicationsTypeModel{
			SiteInfo: siteInfo.Get(),
			Articles: articles.GetType(vars["type"]),
			Type:     vars["type"],
		}

		PublicationsTypeHeaders(c)(w, r)
		err := c.View(w, "pages/publications-type", pageData)
		if err != nil {
			log.Print(err)
		}
	}
}

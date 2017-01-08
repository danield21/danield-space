package handlers

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers"
)

type publicationsModel struct {
	SiteInfo   controllers.SiteInfo
	ArticleMap map[string][]controllers.Article
}

//PublicationsHeaders contains the headers for index
func PublicationsHeaders(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//Publications handles the index page
func Publications(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		siteInfo := controllers.SiteInfoController{}
		articles := controllers.ArticleController{}

		pageData := publicationsModel{
			SiteInfo:   siteInfo.Get(),
			ArticleMap: articles.GetMapKeyedByTypes(5),
		}

		PublicationsTypeHeaders(c)(w, r)
		err := c.View(w, "pages/publications", pageData)
		if err != nil {
			log.Print(err)
		}
	}
}

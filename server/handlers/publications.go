package handlers

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers"
)

type publicationsModel struct {
	SiteInfo     controllers.SiteInfo
	Publications []publicationList
}

type publicationList struct {
	Type     string
	Articles []controllers.Article
}

//PublicationsHeaders contains the headers for index
func PublicationsHeaders(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//Publications handles the index page
func Publications(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		siteInfo := controllers.SiteInfoController{}
		articles := controllers.ArticleController{}

		var publications []publicationList

		for t, a := range articles.GetMapKeyedByTypes(5) {
			publications = append(publications, publicationList{
				Type:     t,
				Articles: a,
			})
		}

		pageData := publicationsModel{
			SiteInfo:     siteInfo.Get(),
			Publications: publications,
		}

		theme := config.GetTheme(r)

		PublicationsTypeHeaders(c)(w, r)
		err := c.View(w, theme, "page/publications", pageData)
		if err != nil {
			log.Print(err)
		}
	}
}

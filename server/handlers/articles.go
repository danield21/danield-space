package handlers

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers"
	"github.com/gorilla/mux"
)

type articlesModel struct {
	SiteInfo controllers.SiteInfo
	Articles []controllers.Article
}

//ArticleHeaders contains the headers for index
func ArticleHeaders(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//Article handles the index page
func Article(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		siteInfo := controllers.SiteInfoController{}
		articles := controllers.ArticleController{}

		pageData := indexModel{
			SiteInfo: siteInfo.Get(),
			Articles: articles.GetType(vars["type"]),
		}

		theme := config.GetTheme(r)

		ArticleHeaders(c)(w, r)
		err := c.View(w, theme, "page/index", pageData)
		if err != nil {
			log.Print(err)
		}
	}
}

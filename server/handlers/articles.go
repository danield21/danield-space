package handlers

import (
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers/articles"
	"github.com/danield21/danield-space/server/controllers/siteInfo"
	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type articlesModel struct {
	SiteInfo siteInfo.SiteInfo
	Articles []articles.Article
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
		context := appengine.NewContext(r)
		vars := mux.Vars(r)

		info, _ := siteInfo.Get(context)

		a, err := articles.GetAllByType(context, vars["type"], 10)
		if err != nil {
			log.Errorf(context, "%v", err)
		}

		pageData := indexModel{
			SiteInfo: info,
			Articles: a,
		}

		theme := config.GetTheme(r)

		ArticleHeaders(c)(w, r)
		err = c.View(w, theme, "page/index", pageData)
		if err != nil {
			log.Errorf(context, "Unable to generate articles page:\n%v", err)
		}
	}
}

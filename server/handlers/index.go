package handlers

import (
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers/articles"
	"github.com/danield21/danield-space/server/controllers/siteInfo"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type indexModel struct {
	SiteInfo siteInfo.SiteInfo
	Articles []articles.Article
}

//IndexHeaders contains the headers for index
func IndexHeaders(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//Index handles the index page
func Index(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := appengine.NewContext(r)

		info, _ := siteInfo.Get(context)

		a, err := articles.GetAll(context, 10)
		if err != nil {
			log.Errorf(context, "%v", err)
		}

		pageData := indexModel{
			SiteInfo: info,
			Articles: a,
		}

		theme := config.GetTheme(r)

		IndexHeaders(c)(w, r)
		err = c.View(w, theme, "page/index", pageData)
		if err != nil {
			log.Errorf(context, "Unable to generate index page:\n%v", err)
		}
	}
}

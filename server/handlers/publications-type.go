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

type publicationsTypeModel struct {
	SiteInfo siteInfo.SiteInfo
	Articles []articles.Article
	Type     string
}

//PublicationsTypeHeaders contains the headers for index
func PublicationsTypeHeaders(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//PublicationsType handles the index page
func PublicationsType(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := appengine.NewContext(r)
		vars := mux.Vars(r)

		info, _ := siteInfo.Get(context)

		a, err := articles.GetAllByType(context, vars["type"], 1)
		if err != nil {
			log.Errorf(context, "%v", err)
		}

		pageData := publicationsTypeModel{
			SiteInfo: info,
			Articles: a,
			Type:     vars["type"],
		}

		theme := config.GetTheme(r)

		PublicationsTypeHeaders(c)(w, r)
		err = c.View(w, theme, "page/publications-type", pageData)
		if err != nil {
			log.Errorf(context, "Unable to generate publications type page:\n%v", err)
		}
	}
}

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

type publicationsModel struct {
	SiteInfo     siteInfo.SiteInfo
	Publications []publicationList
}

type publicationList struct {
	Type     string
	Articles []articles.Article
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
		context := appengine.NewContext(r)

		info, _ := siteInfo.Get(context)

		articleMap, err := articles.GetMapKeyedByTypes(context, 10)
		if err != nil {
			log.Errorf(context, "%v", err)
		}

		var publications []publicationList

		for t, a := range articleMap {
			publications = append(publications, publicationList{
				Type:     t,
				Articles: a,
			})
		}

		pageData := publicationsModel{
			SiteInfo:     info,
			Publications: publications,
		}

		theme := config.GetTheme(r)

		PublicationsTypeHeaders(c)(w, r)
		err = c.View(w, theme, "page/publications", pageData)
		if err != nil {
			log.Errorf(context, "Unable to generate publications page:\n%v", err)
		}
	}
}

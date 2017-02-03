package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers/articles"
	"github.com/danield21/danield-space/server/controllers/siteInfo"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type publishModel struct {
	SiteInfo siteInfo.SiteInfo
	Types    []string
}

//PublishHeaders contains the headers for index
func PublishHeaders(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//Publish handles the index page
func Publish(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		PublishHeaders(c)(w, r)

		context := appengine.NewContext(r)

		info, _ := siteInfo.Get(context)
		types, err := articles.GetTypes(context)
		if err != nil {
			log.Warningf(context, "Unable to get types of articles in Publish handler")
		}

		pageData := publishModel{
			SiteInfo: info,
			Types:    types,
		}

		theme := config.GetTheme(r)
		err = c.View(w, theme, "page/admin/publish", pageData)
		if err != nil {
			log.Errorf(context, "Unable to generate Publish page:\n%v", err)
		}
	}
}

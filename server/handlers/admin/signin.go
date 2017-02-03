package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/config"
	"github.com/danield21/danield-space/server/content"
	"github.com/danield21/danield-space/server/controllers/siteInfo"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type signinModel struct {
	SiteInfo siteInfo.SiteInfo
	Types    []string
}

//SignInHeaders contains the headers for index
func SignInHeaders(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", content.HTML.AddCharset("utf-8").String())
	}
}

//SignIn handles the index page
func SignIn(c config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		PublishHeaders(c)(w, r)

		context := appengine.NewContext(r)

		info, _ := siteInfo.Get(context)

		pageData := signinModel{
			SiteInfo: info,
		}

		theme := config.GetTheme(r)
		err := c.View(w, theme, "page/admin/signin", pageData)
		if err != nil {
			log.Errorf(context, "Unable to generate SignIn page:\n%v", err)
		}
	}
}

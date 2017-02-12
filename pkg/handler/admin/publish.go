package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"github.com/danield21/danield-space/pkg/handler/status"

	"google.golang.org/appengine/log"
)

//PublishHeaders contains the headers for index
func PublishHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Publish handles the index page
func Publish(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))
	session := e.Session(r)

	user, signedIn := GetUser(session)
	if !signedIn {
		status.Unauthorized(e, w, r)
		return
	}

	info := siteInfo.Get(ctx)

	types, err := articles.GetTypes(ctx)
	if err != nil {
		log.Warningf(ctx, "admin.Publish - Unable to get types of articles\n%v", err)
	}

	pageData := struct {
		AdminModel
		Types []string
	}{
		AdminModel: AdminModel{
			BaseModel: handler.BaseModel{
				SiteInfo: info,
			},
			User: user,
		},
		Types: types,
	}

	PublishHeaders(e, w, r)
	err = e.View(w, useTheme, "page/admin/publish", pageData)
	if err != nil {
		log.Errorf(ctx, "admin.Publish - Unable to generate page:\n%v", err)
	}
}

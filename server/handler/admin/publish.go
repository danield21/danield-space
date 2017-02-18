package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"

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

	user, _ := GetUser(session)

	info := siteInfo.Get(ctx)

	cats, err := categories.GetAll(ctx)
	if err != nil {
		log.Warningf(ctx, "admin.Publish - Unable to get types of articles\n%v", err)
	}

	pageData := struct {
		AdminModel
		Categories []*categories.Category
	}{
		AdminModel: AdminModel{
			BaseModel: handler.BaseModel{
				SiteInfo: info,
			},
			User: user,
		},
		Categories: cats,
	}

	PublishHeaders(e, w, r)
	err = e.View(w, useTheme, "page/admin/publish", pageData)
	if err != nil {
		log.Errorf(ctx, "admin.Publish - Unable to generate page:\n%v", err)
	}
}

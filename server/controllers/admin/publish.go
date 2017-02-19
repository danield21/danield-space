package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//PublishHeaders contains the headers for index
func PublishHeaders(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return ctx, nil
}

//Publish handles the index page
func Publish(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := service.Request(ctx)
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
			BaseModel: service.BaseModel{
				SiteInfo: info,
			},
			User: user,
		},
		Categories: cats,
	}

	PublishHeaders(ctx, e, w)
	err = e.View(w, useTheme, "page/admin/publish", pageData)
	if err != nil {
		log.Errorf(ctx, "admin.Publish - Unable to generate page:\n%v", err)
	}
	return ctx, err
}

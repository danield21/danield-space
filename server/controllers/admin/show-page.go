package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//ShowPageHeaders contains the headers
func ShowPageHeaders(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return ctx, nil
}

//ShowPage handles the page
func ShowPage(page string) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		r := service.Request(ctx)
		useTheme := e.Theme(r, theme.GetApp(ctx))
		session := e.Session(r)

		user, _ := GetUser(session)

		info := siteInfo.Get(ctx)

		pageData := struct {
			AdminModel
		}{
			AdminModel: AdminModel{
				BaseModel: service.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
		}

		ShowPageHeaders(ctx, e, w)
		err := e.View(w, useTheme, "page/admin/"+page, pageData)
		if err != nil {
			log.Errorf(ctx, "admin.Index - Unable to generate page:\n%v", err)
		}
		return ctx, err
	}
}

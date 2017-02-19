package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"google.golang.org/appengine/log"
)

//ShowPageHeaders contains the headers
func ShowPageHeaders(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return scp, nil
}

//ShowPage handles the page
func ShowPage(page string) service.Handler {
	return func(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
		r := scp.Request()
		ctx := e.Context(r)
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

		ShowPageHeaders(scp, e, w)
		err := e.View(w, useTheme, "page/admin/"+page, pageData)
		if err != nil {
			log.Errorf(ctx, "admin.Index - Unable to generate page:\n%v", err)
		}
		return scp, err
	}
}

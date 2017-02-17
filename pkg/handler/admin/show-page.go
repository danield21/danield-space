package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"github.com/danield21/danield-space/pkg/repository/siteInfo"
	"github.com/danield21/danield-space/pkg/repository/theme"
	"google.golang.org/appengine/log"
)

//ShowPageHeaders contains the headers
func ShowPageHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//ShowPage handles the page
func ShowPage(page string) handler.Handler {
	return func(e envir.Environment, w http.ResponseWriter, r *http.Request) {
		ctx := e.Context(r)
		useTheme := e.Theme(r, theme.GetApp(ctx))
		session := e.Session(r)

		user, _ := GetUser(session)

		info := siteInfo.Get(ctx)

		pageData := struct {
			AdminModel
		}{
			AdminModel: AdminModel{
				BaseModel: handler.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
		}

		ShowPageHeaders(e, w, r)
		err := e.View(w, useTheme, "page/admin/"+page, pageData)
		if err != nil {
			log.Errorf(ctx, "admin.Index - Unable to generate page:\n%v", err)
		}
	}
}

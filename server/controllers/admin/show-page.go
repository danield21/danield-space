package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//ShowPageHeaders contains the headers
func ShowPageHeaders(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Header().Set("Content-Type", view.HTMLContentType)
	return ctx, nil
}

//ShowPage handles the page
func ShowPage(page string) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		useTheme := "balloon"
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

		ShowPageHeaders(ctx, e, w)
		err := e.View(w, useTheme, "page/admin/"+page, pageData)
		if err != nil {
			log.Errorf(ctx, "admin.Index - Unable to generate page:\n%v", err)
		}
		return ctx, err
	}
}

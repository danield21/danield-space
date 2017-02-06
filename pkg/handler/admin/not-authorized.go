package admin

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
)

//NotAuthorized handles the unauthorized page
func NotAuthorized(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusUnauthorized)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetAdmin(ctx))

	info, _ := siteInfo.Get(ctx)

	pageData := signinModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Redirect: r.URL.Path,
	}

	err := e.View(w, useTheme, "page/admin/not-authorized", pageData)
	if err != nil {
		log.Print(err)
	}
}

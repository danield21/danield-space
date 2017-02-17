package status

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"google.golang.org/appengine/log"
)

//NotFound handles the not found page
func NotFound(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusNotFound)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := handler.BaseModel{
		SiteInfo: info,
	}

	err := e.View(w, useTheme, "page/status/not-found", pageData)
	if err != nil {
		log.Errorf(ctx, "status.NotFound - Unable to generate page\n%v", err)
	}
}

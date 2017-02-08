package status

import (
	"log"
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
)

//NotFound handles the not found page
func NotFound(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusNotFound)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info, _ := siteInfo.Get(ctx)

	pageData := handler.BaseModel{
		SiteInfo: info,
	}

	err := e.View(w, useTheme, "page/status/not-found", pageData)
	if err != nil {
		log.Print(err)
	}
}

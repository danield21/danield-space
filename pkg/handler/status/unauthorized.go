package status

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"google.golang.org/appengine/log"
)

type unauthorizedModel struct {
	handler.BaseModel
	Redirect string
}

//Unauthorized handles the unauthorized page
func Unauthorized(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusUnauthorized)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := unauthorizedModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Redirect: r.URL.Path,
	}

	err := e.View(w, useTheme, "page/status/unauthorized", pageData)
	if err != nil {
		log.Errorf(ctx, "status.Unauthorized - Unable to generate page\n%v", err)
	}
}

package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"google.golang.org/appengine/log"
)

type indexModel struct {
	handler.BaseModel
}

//IndexHeaders contains the headers for index
func IndexHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Index handles the index page
func Index(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetAdmin(ctx))
	session := e.Session(r)

	user, signedIn := session.Values["user"]
	if !signedIn {
		NotAuthorized(e, w, r)
		return
	}
	info, err := siteInfo.Get(ctx)

	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	pageData := indexModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
			User:     user.(string),
		},
	}

	IndexHeaders(e, w, r)
	err = e.View(w, useTheme, "page/admin/index", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate index page:\n%v", err)
	}
}

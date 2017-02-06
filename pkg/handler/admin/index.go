package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"google.golang.org/appengine/log"
)

type indexModel struct {
	SiteInfo siteInfo.SiteInfo
	User     string
}

//IndexHeaders contains the headers for index
func IndexHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Index handles the index page
func Index(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	session := e.Session(r)
	user, signedIn := session.Values["user"]
	if !signedIn {
		NotAuthorized(e, w, r)
		return
	}

	ctx := e.Context(r)
	info, err := siteInfo.Get(ctx)

	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	pageData := indexModel{
		SiteInfo: info,
		User:     user.(string),
	}

	theme := e.Theme(r)

	IndexHeaders(e, w, r)
	err = e.View(w, theme, "page/admin/index", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate index page:\n%v", err)
	}
}

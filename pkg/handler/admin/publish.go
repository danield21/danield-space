package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"

	"google.golang.org/appengine/log"
)

type publishModel struct {
	handler.BaseModel
	Types []string
}

//PublishHeaders contains the headers for index
func PublishHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Publish handles the index page
func Publish(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetAdmin(ctx))
	session := e.Session(r)

	_, signedIn := session.Values["user"]
	if !signedIn {
		NotAuthorized(e, w, r)
		return
	}

	info, _ := siteInfo.Get(ctx)
	types, err := articles.GetTypes(ctx)
	if err != nil {
		log.Warningf(ctx, "Unable to get types of articles in Publish handler")
	}

	pageData := publishModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Types: types,
	}

	PublishHeaders(e, w, r)
	err = e.View(w, useTheme, "page/admin/publish", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate Publish page:\n%v", err)
	}
}

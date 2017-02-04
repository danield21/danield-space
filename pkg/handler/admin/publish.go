package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"

	"google.golang.org/appengine/log"
)

type publishModel struct {
	SiteInfo siteInfo.SiteInfo
	Types    []string
}

//PublishHeaders contains the headers for index
func PublishHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Publish handles the index page
func Publish(e envir.Environment, w http.ResponseWriter, r *http.Request) {

	ctx := e.Context(r)

	info, _ := siteInfo.Get(ctx)
	types, err := articles.GetTypes(ctx)
	if err != nil {
		log.Warningf(ctx, "Unable to get types of articles in Publish handler")
	}

	pageData := publishModel{
		SiteInfo: info,
		Types:    types,
	}

	PublishHeaders(e, w, r)
	theme := e.Theme(r)
	err = e.View(w, theme, "page/admin/publish", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate Publish page:\n%v", err)
	}
}

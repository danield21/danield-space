package app

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type publicationsTypeModel struct {
	SiteInfo siteInfo.SiteInfo
	Articles []articles.Article
	Type     string
}

//PublicationsTypeHeaders contains the headers for index
func PublicationsTypeHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//PublicationsType handles the index page
func PublicationsType(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	vars := mux.Vars(r)

	info, _ := siteInfo.Get(ctx)

	a, err := articles.GetAllByType(ctx, vars["type"], 1)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	pageData := publicationsTypeModel{
		SiteInfo: info,
		Articles: a,
		Type:     vars["type"],
	}

	theme := e.Theme(r)

	PublicationsTypeHeaders(e, w, r)
	err = e.View(w, theme, "page/app/publications-type", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate publications type page:\n%v", err)
	}
}

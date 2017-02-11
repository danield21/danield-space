package app

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type publicationsTypeModel struct {
	handler.BaseModel
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
	useTheme := e.Theme(r, theme.GetApp(ctx))
	vars := mux.Vars(r)

	info, err := siteInfo.Get(ctx)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to get site information\n%v", err)
	}

	a, err := articles.GetAllByType(ctx, vars["type"], 1)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to get articles by type\n%v", err)
	}

	pageData := publicationsTypeModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Articles: a,
		Type:     vars["type"],
	}

	PublicationsTypeHeaders(e, w, r)
	err = e.View(w, useTheme, "page/app/publications-type", pageData)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to generate page\n%v", err)
	}
}

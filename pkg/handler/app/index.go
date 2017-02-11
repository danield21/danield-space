package app

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"google.golang.org/appengine/log"
)

type indexModel struct {
	handler.BaseModel
	Articles []articles.Article
}

//IndexHeaders contains the headers for index
func IndexHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Index handles the index page
func Index(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info, err := siteInfo.Get(ctx)
	if err != nil {
		log.Errorf(ctx, "app.Index - Unable to get site information\n%v", err)
	}

	a, err := articles.GetAll(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "app.Index - Unable to get last 10 articles\n%v", err)
	}

	pageData := indexModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Articles: a,
	}

	IndexHeaders(e, w, r)
	err = e.View(w, useTheme, "page/app/index", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Index - Unable to generate page\n%v", err)
	}
}

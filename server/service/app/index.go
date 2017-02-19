package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"google.golang.org/appengine/log"
)

type indexModel struct {
	service.BaseModel
	Articles []*articles.Article
}

//IndexHeaders contains the headers for index
func IndexHeaders(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return scp, nil
}

//Index handles the index page
func Index(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	a, err := articles.GetAll(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "app.Index - Unable to get last 10 articles\n%v", err)
	}

	pageData := indexModel{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Articles: a,
	}

	IndexHeaders(scp, e, w)
	err = e.View(w, useTheme, "page/app/index", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Index - Unable to generate page\n%v", err)
	}
	return scp, err
}

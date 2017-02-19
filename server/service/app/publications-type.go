package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type publicationsTypeModel struct {
	service.BaseModel
	Articles []*articles.Article
	Category *categories.Category
}

//PublicationsTypeHeaders contains the headers for index
func PublicationsTypeHeaders(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return nil
}

//PublicationsType handles the index page
func PublicationsType(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
	r := scp.Request()
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))
	vars := mux.Vars(r)

	info := siteInfo.Get(ctx)

	cat, err := categories.Get(ctx, vars["category"])
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to get category\n%v", err)
	}

	a, err := articles.GetAllByCategory(ctx, cat, 1)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to get articles by category\n%v", err)
	}

	pageData := publicationsTypeModel{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Articles: a,
		Category: cat,
	}

	PublicationsTypeHeaders(scp, e, w)
	err = e.View(w, useTheme, "page/app/publications-type", pageData)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to generate page\n%v", err)
	}
	return err
}

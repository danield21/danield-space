package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type publicationsTypeModel struct {
	handler.BaseModel
	Articles []articles.Article
	Category categories.Category
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

	info := siteInfo.Get(ctx)

	cat, _, err := categories.Get(ctx, vars["category"])
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to get category\n%v", err)
	}

	a, err := articles.GetAllByCategory(ctx, vars["category"], 1)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to get articles by category\n%v", err)
	}

	pageData := publicationsTypeModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Articles: a,
		Category: cat,
	}

	PublicationsTypeHeaders(e, w, r)
	err = e.View(w, useTheme, "page/app/publications-type", pageData)
	if err != nil {
		log.Errorf(ctx, "app.PublicationsType - Unable to generate page\n%v", err)
	}
}

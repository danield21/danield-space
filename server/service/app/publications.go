package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"google.golang.org/appengine/log"
)

type publicationsModel struct {
	service.BaseModel
	Publications []publicationList
}

type publicationList struct {
	Category *categories.Category
	Articles []*articles.Article
}

//PublicationsHeaders contains the headers for index
func PublicationsHeaders(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return scp, nil
}

//Publications handles the index page
func Publications(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	articleMap, err := articles.GetMapKeyedByCategory(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "app.Publications - Unable to get articles organized by their type\n%v", err)
	}

	var publications []publicationList

	for cat, a := range articleMap {
		publications = append(publications, publicationList{
			Category: cat,
			Articles: a,
		})
	}

	pageData := publicationsModel{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Publications: publications,
	}

	PublicationsTypeHeaders(scp, e, w)
	err = e.View(w, useTheme, "page/app/publications", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Publications - Unable to generate page\n%v", err)
	}
	return scp, err
}

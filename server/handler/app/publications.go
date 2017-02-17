package app

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"google.golang.org/appengine/log"
)

type publicationsModel struct {
	handler.BaseModel
	Publications []publicationList
}

type publicationList struct {
	Type     string
	Articles []articles.Article
}

//PublicationsHeaders contains the headers for index
func PublicationsHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Publications handles the index page
func Publications(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	articleMap, err := articles.GetMapKeyedByTypes(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "app.Publications - Unable to get articles organized by their type\n%v", err)
	}

	var publications []publicationList

	for t, a := range articleMap {
		publications = append(publications, publicationList{
			Type:     t,
			Articles: a,
		})
	}

	pageData := publicationsModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Publications: publications,
	}

	PublicationsTypeHeaders(e, w, r)
	err = e.View(w, useTheme, "page/app/publications", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Publications - Unable to generate page\n%v", err)
	}
}

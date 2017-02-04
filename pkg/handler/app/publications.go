package app

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"
	"google.golang.org/appengine/log"
)

type publicationsModel struct {
	SiteInfo     siteInfo.SiteInfo
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
	info, _ := siteInfo.Get(ctx)

	articleMap, err := articles.GetMapKeyedByTypes(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	var publications []publicationList

	for t, a := range articleMap {
		publications = append(publications, publicationList{
			Type:     t,
			Articles: a,
		})
	}

	pageData := publicationsModel{
		SiteInfo:     info,
		Publications: publications,
	}

	theme := e.Theme(r)

	PublicationsTypeHeaders(e, w, r)
	err = e.View(w, theme, "page/publications", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate publications page:\n%v", err)
	}
}

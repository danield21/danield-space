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

type articlesModel struct {
	SiteInfo siteInfo.SiteInfo
	Articles []articles.Article
}

//ArticleHeaders contains the headers for index
func ArticleHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Article handles the index page
func Article(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	vars := mux.Vars(r)

	info, _ := siteInfo.Get(ctx)

	a, err := articles.GetAllByType(ctx, vars["type"], 10)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	pageData := indexModel{
		SiteInfo: info,
		Articles: a,
	}

	theme := e.Theme(r)

	ArticleHeaders(e, w, r)
	err = e.View(w, theme, "page/index", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate articles page:\n%v", err)
	}
}

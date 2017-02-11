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

type articlesModel struct {
	handler.BaseModel
	Articles []articles.Article
}

//ArticleHeaders contains the headers for index
func ArticleHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//Article handles the index page
func Article(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))
	vars := mux.Vars(r)

	info := siteInfo.Get(ctx)

	a, err := articles.GetAllByType(ctx, vars["type"], 10)
	if err != nil {
		log.Errorf(ctx, "app.Article - Unable to get articles by type\n%v", err)
	}

	pageData := articlesModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Articles: a,
	}

	ArticleHeaders(e, w, r)
	err = e.View(w, useTheme, "page/app/index", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Article - Unable to generate page:\n%v", err)
	}
}

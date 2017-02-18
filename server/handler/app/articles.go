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

type articlesModel struct {
	handler.BaseModel
	Article *articles.Article
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
	cat := categories.EmptyCategory(vars["category"])

	a, err := articles.Get(ctx, cat, vars["key"])
	if err != nil {
		log.Errorf(ctx, "app.Article - Unable to get articles by type\n%v", err)
	}

	pageData := articlesModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Article: a,
	}

	ArticleHeaders(e, w, r)
	err = e.View(w, useTheme, "page/app/article", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Article - Unable to generate page:\n%v", err)
	}
}

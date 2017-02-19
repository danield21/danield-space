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
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type articlesModel struct {
	service.BaseModel
	Article *articles.Article
}

//ArticleHeaders contains the headers for index
func ArticleHeaders(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return ctx, nil
}

//Article handles the index page
func Article(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := service.Request(ctx)
	useTheme := e.Theme(r, theme.GetApp(ctx))
	vars := mux.Vars(r)

	info := siteInfo.Get(ctx)
	cat := categories.EmptyCategory(vars["category"])

	a, err := articles.Get(ctx, cat, vars["key"])
	if err != nil {
		log.Errorf(ctx, "app.Article - Unable to get articles by type\n%v", err)
	}

	pageData := articlesModel{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Article: a,
	}

	ArticleHeaders(ctx, e, w)
	err = e.View(w, useTheme, "page/app/article", pageData)
	if err != nil {
		log.Errorf(ctx, "app.Article - Unable to generate page:\n%v", err)
	}
	return ctx, err
}

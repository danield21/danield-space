package controllers

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/process"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
)

//AdminRouter creates a new server instance to run
type AdminRouter struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Account      store.AccountRepository
	About        store.AboutRepository
	Article      store.ArticleRepository
	Category     store.CategoryRepository
	NotFound     http.Handler
	Unauthorized http.Handler
}

func (rtr AdminRouter) Route(r *mux.Router) {
	r.NotFoundHandler = rtr.NotFound

	r.Handle("/", admin.IndexHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Article:      rtr.Article,
		Category:     rtr.Category,
		Unauthorized: rtr.Unauthorized,
	})
	r.Handle("/sign-in", admin.SignInHandler{
		Context:  rtr.Context,
		Session:  rtr.Session,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
		Account:  rtr.Account,
		SignIn: process.SignInProcessor{
			Account: rtr.Account,
		},
	})
	r.Handle("/sign-out", admin.SignOutHandler{
		Context:  rtr.Context,
		Session:  rtr.Session,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
		SignOut:  process.SignOutProcessor,
	})
	r.Handle("/about", admin.AboutHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		About:        rtr.About,
		Unauthorized: rtr.Unauthorized,
		PutAbout: process.PutAboutProcessor{
			About: rtr.About,
		},
	})
	r.Handle("/account", admin.AccountAllHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Account:      rtr.Account,
		Unauthorized: rtr.Unauthorized,
	})
	r.Handle("/account/create", admin.AccountCreateHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Account:      rtr.Account,
		Unauthorized: rtr.Unauthorized,
		PutAccount: process.PutAccountProcessor{
			Account: rtr.Account,
		},
	})
	r.Handle("/article", admin.ArticleAllHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Article:      rtr.Article,
		Unauthorized: rtr.Unauthorized,
	})
	r.Handle("/article/publish", admin.ArticlePublishHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Category:     rtr.Category,
		Unauthorized: rtr.Unauthorized,
		PutArticle: process.PutArticleProcessor{
			Article:  rtr.Article,
			Category: rtr.Category,
		},
	})
	r.Handle("/category", admin.CategoryAllHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Category:     rtr.Category,
		Unauthorized: rtr.Unauthorized,
	})
	r.Handle("/category/create", admin.CategoryCreateHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Unauthorized: rtr.Unauthorized,
		PutCategory: process.PutCategoryProcessor{
			Category: rtr.Category,
		},
	})
	r.Handle("/site-info", admin.SiteInfoHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Unauthorized: rtr.Unauthorized,
		PutSiteInfo: process.PutSiteInfoProcessor{
			SiteInfo: rtr.SiteInfo,
		},
	})
}

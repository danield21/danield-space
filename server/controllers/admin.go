package controllers

import (
	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/process"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
)

//AdminRouter creates a new server instance to run
type AdminRouter struct {
	Context  handler.ContextGenerator
	Session  handler.SessionGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	Account  store.AccountRepository
	About    store.AboutRepository
	Article  store.ArticleRepository
	Category store.CategoryRepository
}

func (rtr AdminRouter) Route(r *mux.Router) {
	notFnd := status.NotFoundHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
	}

	unauth := status.UnauthorizedHandler{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
	}

	r.NotFoundHandler = notFnd

	r.Handle("/", admin.IndexHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Article:      rtr.Article,
		Category:     rtr.Category,
		Unauthorized: unauth,
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
		Unauthorized: unauth,
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
		Unauthorized: unauth,
	})
	r.Handle("/account/create", admin.AccountCreateHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Account:      rtr.Account,
		Unauthorized: unauth,
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
		Unauthorized: unauth,
	})
	r.Handle("/article/publish", admin.ArticlePublishHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Category:     rtr.Category,
		Unauthorized: unauth,
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
		Unauthorized: unauth,
	})
	r.Handle("/category/create", admin.CategoryCreateHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Unauthorized: unauth,
		PutCategory: process.PutCategoryProcessor{
			Category: rtr.Category,
		},
	})
	r.Handle("/site-info", admin.SiteInfoHandler{
		Context:      rtr.Context,
		Session:      rtr.Session,
		Renderer:     rtr.Renderer,
		SiteInfo:     rtr.SiteInfo,
		Unauthorized: unauth,
		PutSiteInfo: process.PutSiteInfoProcessor{
			SiteInfo: rtr.SiteInfo,
		},
	})
}

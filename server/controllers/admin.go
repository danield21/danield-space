package controllers

import (
	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/danield21/danield-space/server/controllers/process"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/mux"
)

//Admin configures the handlers for admin handlers
func Admin(e handler.Environment, r *mux.Router) {

	mgr := Migrator{
		Environment: e,
	}

	Unauthorized := status.UnauthorizedHandler{
		Context:  mgr.Context(),
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
	}
	NotFound := status.NotFoundHandler{
		Context:  mgr.Context(),
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
	}

	r.NotFoundHandler = NotFound

	r.Handle("/", admin.IndexHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Article:      e.Repository().Article(),
		Category:     e.Repository().Category(),
		Unauthorized: Unauthorized,
	})
	r.Handle("/sign-in", admin.SignInHandler{
		Context:  mgr.Context(),
		Session:  mgr.Session(),
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
		Account:  e.Repository().Account(),
		SignIn: process.SignInProcessor{
			Account: e.Repository().Account(),
		},
	})
	r.Handle("/sign-out", admin.SignOutHandler{
		Context:  mgr.Context(),
		Session:  mgr.Session(),
		Renderer: mgr,
		SiteInfo: e.Repository().SiteInfo(),
		SignOut:  process.SignOutProcessor,
	})
	r.Handle("/about", admin.AboutHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		About:        e.Repository().About(),
		Unauthorized: Unauthorized,
		PutAbout: process.PutAboutProcessor{
			About: e.Repository().About(),
		},
	})
	r.Handle("/account", admin.AccountAllHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Account:      e.Repository().Account(),
		Unauthorized: Unauthorized,
	})
	r.Handle("/account/create", admin.AccountCreateHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Account:      e.Repository().Account(),
		Unauthorized: Unauthorized,
		PutAccount: process.PutAccountProcessor{
			Account: e.Repository().Account(),
		},
	})
	r.Handle("/article", admin.ArticleAllHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Article:      e.Repository().Article(),
		Unauthorized: Unauthorized,
	})
	r.Handle("/article/publish", admin.ArticlePublishHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Category:     e.Repository().Category(),
		Unauthorized: Unauthorized,
		PutArticle: process.PutArticleProcessor{
			Article:  e.Repository().Article(),
			Category: e.Repository().Category(),
		},
	})
	r.Handle("/category", admin.CategoryAllHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Category:     e.Repository().Category(),
		Unauthorized: Unauthorized,
	})
	r.Handle("/category/create", admin.CategoryCreateHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Unauthorized: Unauthorized,
		PutCategory: process.PutCategoryProcessor{
			Category: e.Repository().Category(),
		},
	})
	r.Handle("/site-info", admin.SiteInfoHandler{
		Context:      mgr.Context(),
		Session:      mgr.Session(),
		Renderer:     mgr,
		SiteInfo:     e.Repository().SiteInfo(),
		Unauthorized: Unauthorized,
		PutSiteInfo: process.PutSiteInfoProcessor{
			SiteInfo: e.Repository().SiteInfo(),
		},
	})
}

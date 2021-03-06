package router

import (
	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/controller/admin"
	"github.com/danield21/danield-space/server/controller/status"
	"github.com/danield21/danield-space/server/process"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
)

//AdminRouter creates a new server instance to run
type AdminRouter struct {
	Context  ContextGenerator
	Session  SessionGenerator
	Renderer controller.Renderer
	SiteInfo store.SiteInfoRepository
	About    store.AboutRepository
	Article  store.ArticleRepository
	Category store.CategoryRepository
}

func (rtr AdminRouter) Route(r *mux.Router) {
	ctrHnd := controller.ControllerHandler{
		Renderer: rtr.Renderer,
		Session:  rtr.Session,
		Context:  rtr.Context,
	}

	inrErr := status.InternalServerErrorController{
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
	}

	notFnd := status.NotFoundController{
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
	}

	unauth := status.UnauthorizedController{
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
	}

	r.NotFoundHandler = ctrHnd.ToHandler(notFnd)

	r.Handle("/", ctrHnd.ToHandler(admin.IndexController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Article:             rtr.Article,
		Category:            rtr.Category,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
	}))
	r.Handle("/about", ctrHnd.ToHandler(admin.AboutController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		About:               rtr.About,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
		PutAbout: process.PutAboutProcessor{
			About: rtr.About,
		},
	}))
	r.Handle("/article", ctrHnd.ToHandler(admin.ArticleAllController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Article:             rtr.Article,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
	}))
	r.Handle("/article/publish", ctrHnd.ToHandler(admin.ArticlePublishController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Category:            rtr.Category,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
		PutArticle: process.PutArticleProcessor{
			Article:  rtr.Article,
			Category: rtr.Category,
		},
	}))
	r.Handle("/category", ctrHnd.ToHandler(admin.CategoryAllController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Category:            rtr.Category,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
	}))
	r.Handle("/category/create", ctrHnd.ToHandler(admin.CategoryCreateController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
		PutCategory: process.PutCategoryProcessor{
			Category: rtr.Category,
		},
	}))
	r.Handle("/site-info", ctrHnd.ToHandler(admin.SiteInfoController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Unauthorized:        unauth,
		InternalServerError: inrErr,
		PutSiteInfo: process.PutSiteInfoProcessor{
			SiteInfo: rtr.SiteInfo,
		},
	}))
}

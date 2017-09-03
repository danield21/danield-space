package controllers

import (
	"github.com/danield21/danield-space/server/controllers/app"
	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/mux"
)

//App creates a new server instance to run
type AppRouter struct {
	Context  ContextGenerator
	Session  SessionGenerator
	Renderer Renderer
	SiteInfo store.SiteInfoRepository
	About    store.AboutRepository
	Article  store.ArticleRepository
	Category store.CategoryRepository
}

func (rtr AppRouter) Route(r *mux.Router) {
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
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		InternalServerError: inrErr,
	}

	r.NotFoundHandler = ctrHnd.ToHandler(notFnd)

	r.Handle("/", ctrHnd.ToHandler(app.IndexController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Article:             rtr.Article,
		InternalServerError: inrErr,
	}))
	r.Handle("/about", ctrHnd.ToHandler(app.AboutController{
		Context:  rtr.Context,
		Renderer: rtr.Renderer,
		SiteInfo: rtr.SiteInfo,
		About:    rtr.About,
	}))
	r.Handle("/articles", ctrHnd.ToHandler(app.ArticlesController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Article:             rtr.Article,
		InternalServerError: inrErr,
	}))
	r.Handle("/articles/{category}", ctrHnd.ToHandler(app.ArticleCategoryController{
		Renderer:            rtr.Renderer,
		NotFound:            notFnd,
		SiteInfo:            rtr.SiteInfo,
		Article:             rtr.Article,
		Category:            rtr.Category,
		InternalServerError: inrErr,
	}))
	r.Handle("/articles/{category}/{key}", ctrHnd.ToHandler(app.ArticleController{
		Renderer:            rtr.Renderer,
		SiteInfo:            rtr.SiteInfo,
		Article:             rtr.Article,
		NotFound:            notFnd,
		InternalServerError: inrErr,
	}))
}

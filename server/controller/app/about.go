package app

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

type AboutController struct {
	Context             controller.ContextGenerator
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	About               store.AboutRepository
	InternalServerError controller.Controller
}

func (ctr AboutController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	abt, err := ctr.About.Get(ctx)

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to get about contents"))
		return ctr.InternalServerError
	}

	cnt, err := ctr.Renderer.String("page/app/about", struct {
		About template.HTML
	}{
		abt,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

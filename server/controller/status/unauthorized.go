package status

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type UnauthorizedController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	About               store.AboutRepository
	InternalServerError controller.Controller
}

func (ctr UnauthorizedController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	cnt, err := ctr.Renderer.String("page/status/unauthorized", nil)

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	pg.Status = http.StatusUnauthorized
	pg.Content = template.HTML(cnt)

	return nil
}

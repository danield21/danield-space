package status

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type NotFoundController struct {
	Renderer            Renderer
	SiteInfo            store.SiteInfoRepository
	InternalServerError controller.Controller
}

func (ctr NotFoundController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	cnt, err := ctr.Renderer.String("page/status/not-found", nil)

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	pg.Status = http.StatusNotFound
	pg.Content = template.HTML(cnt)

	return nil
}

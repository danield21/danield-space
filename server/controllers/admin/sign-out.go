package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type SignOutController struct {
	Renderer            handler.Renderer
	SiteInfo            store.SiteInfoRepository
	SignOut             handler.Processor
	InternalServerError controller.Controller
}

func (ctr SignOutController) Serve(ctx context.Context, pg *handler.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	if rqs.Method == http.MethodPost {
		ctr.SignOut.Process(ctx, rqs, pg.Session)
	}

	cnt, err := ctr.Renderer.Render(ctx, "page/admin/sign-out", nil)

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Status = http.StatusSeeOther
	pg.Header["Location"] = "/"
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

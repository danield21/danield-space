package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/controllers/process"
	"github.com/danield21/danield-space/server/controllers/session"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type SiteInfoController struct {
	Renderer            Renderer
	SiteInfo            store.SiteInfoRepository
	About               store.AboutRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
	PutSiteInfo         Processor
}

func (ctr SiteInfoController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := session.User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	frm := form.NewForm()

	if rqs.Method == http.MethodPost {
		frm = ctr.PutSiteInfo.Process(ctx, rqs, pg.Session)
	}

	if frm.IsEmpty() {
		frm = process.RepackSiteInfo(info)
	}

	cnt, err := ctr.Renderer.String("page/admin/site-info-manage", struct {
		User string
		Form form.Form
	}{
		User: usr,
		Form: frm,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Content = template.HTML(cnt)

	return nil
}

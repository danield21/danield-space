package admin

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/process"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type SiteInfoController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	About               store.AboutRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
	PutSiteInfo         Processor
}

func (ctr SiteInfoController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	usr := user.Current(ctx)
	if usr == nil {
		return ctr.Unauthorized
	}

	signOut, err := user.LogoutURL(ctx, "/")
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "cannot create a url for logging out"))
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
		User    string
		Form    form.Form
		SignOut string
	}{
		User:    usr.String(),
		Form:    frm,
		SignOut: signOut,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Content = template.HTML(cnt)

	return nil
}

package admin

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

type AboutController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	About               store.AboutRepository
	InternalServerError controller.Controller
	Unauthorized        controller.Controller
	PutAbout            Processor
}

func (ctr AboutController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := User(pg.Session)
	if !signedIn {
		log.Errorf(ctx, "%v", errors.New("user is not authenicated"))
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	frm := form.NewForm()

	if rqs.Method == http.MethodPost {
		frm = ctr.PutAbout.Process(ctx, rqs, pg.Session)
	}

	if frm.IsEmpty() {
		html, err := ctr.About.Get(ctx)
		if err == nil {
			abtFld := new(form.Field)
			abtFld.Values = []string{string(html)}
			frm.Fields["about"] = abtFld
		} else {
			log.Warningf(ctx, "Unable to get about summary\n%v", err)
		}
	}

	cnt, err := ctr.Renderer.String("page/admin/about", struct {
		User string
		Form form.Form
	}{
		User: usr,
		Form: frm,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

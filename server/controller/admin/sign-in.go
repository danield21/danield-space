package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type SignInController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Account             store.AccountRepository
	SignIn              Processor
	InternalServerError controller.Controller
}

func (ctr SignInController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	var frm form.Form

	if rqs.Method == http.MethodPost {
		frm = ctr.SignIn.Process(ctx, rqs, pg.Session)
	}

	cnt, err := ctr.Renderer.String("page/admin/sign-in", struct {
		Form form.Form
	}{
		frm,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	if frm.IsSuccessful() {
		pg.Status = http.StatusSeeOther
		pg.Header["Location"] = "."
	}
	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

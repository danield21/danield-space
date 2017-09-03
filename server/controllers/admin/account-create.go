package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/controllers/process"
	"github.com/danield21/danield-space/server/controllers/session"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type AccountCreateController struct {
	Renderer            handler.Renderer
	SiteInfo            store.SiteInfoRepository
	Account             store.AccountRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
	PutAccount          handler.Processor
}

func (ctr AccountCreateController) Serve(ctx context.Context, pg *handler.Page, rqs *http.Request) controller.Controller {

	rqs.ParseForm()

	usr, signedIn := session.User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	current, err := ctr.Account.Get(ctx, usr)
	if err != nil {
		log.Warningf(ctx, "admin.AccountAllHandler - Unable to verify account %s\n%v", usr, err)
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	frm := form.NewForm()

	if rqs.Method == http.MethodPost {
		frm = ctr.PutAccount.Process(ctx, rqs, pg.Session)
	}

	target := rqs.Form.Get("account")
	if frm.IsEmpty() && target != "" {
		tUser, err := ctr.Account.Get(ctx, target)
		if err == nil {
			frm = process.AccountToForm(tUser)
		} else {
			log.Warningf(ctx, "Unable to get account %s\n%v", target, err)
		}
	}

	cnt, err := ctr.Renderer.Render(ctx, "page/admin/account-create", struct {
		User  string
		Form  form.Form
		Super bool
	}{
		User:  usr,
		Form:  frm,
		Super: current.Super,
	})

	if err != nil {
		log.Errorf(ctx, "admin.AccountAllHandler - Unable to render content\n%v", err)
		return ctr.InternalServerError
	}

	pg.Content = template.HTML(cnt)

	return nil
}

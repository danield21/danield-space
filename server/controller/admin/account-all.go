package admin

import (
	"context"
	"errors"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type AccountAllController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Account             store.AccountRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr AccountAllController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := User(pg.Session)
	if !signedIn {
		log.Errorf(ctx, "%v", errors.New("user is not authenicated"))
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

	acts, err := ctr.Account.GetAll(ctx)
	if err != nil {
		log.Errorf(ctx, "admin.AccountAllHandler - Unable to get all acounts\n%v", err)
	}

	cnt, err := ctr.Renderer.String("page/admin/account-all", struct {
		User     string
		Accounts []*store.Account
		Super    bool
	}{
		User:     usr,
		Accounts: acts,
		Super:    current.Super,
	})

	if err != nil {
		log.Errorf(ctx, "admin.AccountAllHandler - Unable to render content\n%v", err)
		return ctr.InternalServerError
	}

	pg.Content = template.HTML(cnt)

	return nil
}

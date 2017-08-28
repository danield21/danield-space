package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type AccountAllHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Account      store.AccountRepository
	Unauthorized http.Handler
}

func (hnd AccountAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	usr, signedIn := link.User(ses)
	if !signedIn {
		hnd.Unauthorized.ServeHTTP(w, r)
		return
	}

	current, err := hnd.Account.Get(ctx, usr)
	if err != nil {
		log.Warningf(ctx, "admin.AccountAllHandler - Unable to verify account %s\n%v", usr, err)
		hnd.Unauthorized.ServeHTTP(w, r)
		return
	}

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	acts, err := hnd.Account.GetAll(ctx)
	if err != nil {
		log.Errorf(ctx, "admin.AccountAllHandler - Unable to get all acounts\n%v", err)
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/account-all", struct {
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
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

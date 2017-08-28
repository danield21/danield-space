package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/process"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type AccountCreateHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Account      store.AccountRepository
	Unauthorized http.Handler
	PutAccount   handler.Processor
}

func (hnd AccountCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()
	r.ParseForm()

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

	frm := form.NewForm()

	if r.Method == http.MethodPost {
		frm = hnd.PutAccount.Process(ctx, r, ses)
	}

	target := r.Form.Get("account")
	if frm.IsEmpty() && target != "" {
		tUser, err := hnd.Account.Get(ctx, target)
		if err == nil {
			frm = process.AccountToForm(tUser)
		} else {
			log.Warningf(ctx, "Unable to get account %s\n%v", target, err)
		}
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/account-create", struct {
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
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

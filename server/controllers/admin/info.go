package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/action"
	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type SiteInfoHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	About        store.AboutRepository
	Unauthorized http.Handler
	PutSiteInfo  handler.Processor
}

func (hnd SiteInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	usr, signedIn := link.User(ses)
	if !signedIn {
		hnd.Unauthorized.ServeHTTP(w, r)
		return
	}

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	frm := form.NewForm()

	if r.Method == http.MethodPost {
		frm = hnd.PutSiteInfo.Process(ctx, r, ses)
	}

	if frm.IsEmpty() {
		frm = action.RepackSiteInfo(info)
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/site-info-manage", struct {
		User string
		Form form.Form
	}{
		User: usr,
		Form: frm,
	})

	if err != nil {
		log.Errorf(ctx, "admin.IndexHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

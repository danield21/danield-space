package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/session"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type AboutHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	About        store.AboutRepository
	Unauthorized http.Handler
	PutAbout     handler.Processor
}

func (hnd AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	usr, signedIn := session.User(ses)
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
		frm = hnd.PutAbout.Process(ctx, r, ses)
	}

	if frm.IsEmpty() {
		html, err := hnd.About.Get(ctx)
		if err == nil {
			abtFld := new(form.Field)
			abtFld.Values = []string{string(html)}
			frm.Fields["about"] = abtFld
		} else {
			log.Warningf(ctx, "Unable to get about summary\n%v", err)
		}
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/about", struct {
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

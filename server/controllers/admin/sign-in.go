package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type SignInHandler struct {
	Context  handler.ContextGenerator
	Session  handler.SessionGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	Account  store.AccountRepository
	SignIn   handler.Processor
}

func (hnd SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	var frm form.Form

	if r.Method == http.MethodPost {
		frm = hnd.SignIn.Process(ctx, r, ses)
	}

	if frm.IsSuccessful() {
		pg.Status = http.StatusSeeOther
		pg.Header["Location"] = "."
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/sign-in", struct {
		Form form.Form
	}{
		frm,
	})

	if err != nil {
		log.Errorf(ctx, "admin.SignInHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

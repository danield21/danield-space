package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
)

type SignOutHandler struct {
	Context  handler.ContextGenerator
	Session  handler.SessionGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	SignOut  handler.Processor
}

func (hnd SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	if r.Method == http.MethodPost {
		hnd.SignOut.Process(ctx, r, ses)
	}

	pg.Status = http.StatusSeeOther
	pg.Header["Location"] = "/"

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/sign-out", nil)

	if err != nil {
		log.Errorf(ctx, "admin.SignInHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

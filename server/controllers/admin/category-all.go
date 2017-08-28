package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
)

type CategoryAllHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Category     store.CategoryRepository
	Unauthorized http.Handler
}

func (hnd CategoryAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	cats, err := hnd.Category.GetAll(ctx)
	if err != nil {
		log.Errorf(ctx, "admin.CategoryAllHandler - Unable to get all categories\n%v", err)
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/category-all", struct {
		User       string
		Categories []*store.Category
	}{
		User:       usr,
		Categories: cats,
	})

	if err != nil {
		log.Errorf(ctx, "admin.CategoryAllHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

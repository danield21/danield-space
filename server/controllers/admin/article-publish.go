package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/form"

	"github.com/danield21/danield-space/server/controllers/session"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type ArticlePublishHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Category     store.CategoryRepository
	Unauthorized http.Handler
	PutArticle   handler.Processor
}

func (hnd ArticlePublishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		frm = hnd.PutArticle.Process(ctx, r, ses)
	}

	cats, err := hnd.Category.GetAll(ctx)
	if err != nil {
		log.Warningf(ctx, "admin.Publish - Unable to get types of articles\n%v", err)
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/article-publish", struct {
		User       string
		Form       form.Form
		Categories []*store.Category
	}{
		User:       usr,
		Form:       frm,
		Categories: cats,
	})

	if err != nil {
		log.Errorf(ctx, "admin.ArticleAllHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
)

type ArticleAllHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Article      store.ArticleRepository
	Unauthorized http.Handler
}

func (hnd ArticleAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	arts, err := hnd.Article.GetAll(ctx, -1)
	if err != nil {
		log.Errorf(ctx, "admin.ArticleAllHandler - Unable to get all acounts\n%v", err)
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/article-all", struct {
		User     string
		Articles []*store.Article
	}{
		User:     usr,
		Articles: arts,
	})

	if err != nil {
		log.Errorf(ctx, "admin.ArticleAllHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

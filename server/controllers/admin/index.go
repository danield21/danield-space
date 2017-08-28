package admin

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
)

type IndexHandler struct {
	Context      handler.ContextGenerator
	Session      handler.SessionGenerator
	Renderer     handler.Renderer
	SiteInfo     store.SiteInfoRepository
	Article      store.ArticleRepository
	Category     store.CategoryRepository
	Unauthorized http.Handler
}

func (hnd IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	user, signedIn := link.User(ses)
	if !signedIn {
		hnd.Unauthorized.ServeHTTP(w, r)
		return
	}

	info := hnd.SiteInfo.Get(ctx)
	cats, _ := hnd.Category.GetAll(ctx)
	arts, _ := hnd.Article.GetAll(ctx, 1)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/index", struct {
		User          string
		HasCategories bool
		HasArticles   bool
	}{
		User:          user,
		HasCategories: len(cats) > 0,
		HasArticles:   len(arts) > 0,
	})

	if err != nil {
		log.Errorf(ctx, "admin.IndexHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

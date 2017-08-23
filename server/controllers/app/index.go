package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type IndexHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	Article  store.ArticleRepository
}

func (hnd IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.New(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	a, err := hnd.Article.GetAll(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "app.IndexPageLink - Unable to get last 10 articles\n%v", err)
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/app/index", struct {
		Articles    []*store.Article
		Description string
	}{
		Articles:    a,
		Description: info.Description,
	})

	if err != nil {
		log.Errorf(ctx, "app.AboutHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	hnd.Renderer.Send(w, r, pg)
}

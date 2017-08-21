package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type AboutHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	About    store.AboutRepository
}

func (hnd AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.New(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	abt, err := hnd.About.Get(ctx)

	if err != nil {
		log.Errorf(ctx, "app.AboutHandler - Unable to get about contents\n%v", err)
		return
	}

	pg.Content = template.HTML(hnd.Renderer.Render(ctx, "page/app/about", struct {
		About template.HTML
	}{
		abt,
	}))

	hnd.Renderer.Send(w, r, pg)
}

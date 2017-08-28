package status

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/log"
)

type NotFoundHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	About    store.AboutRepository
}

func (hnd NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.Generate(r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	cnt, err := hnd.Renderer.Render(ctx, "page/status/not-found", nil)

	if err != nil {
		log.Errorf(ctx, "status.NotFound - Unable to render content\n%v", err)
		return
	}

	pg.Status = http.StatusNotFound
	pg.Content = template.HTML(cnt)

	hnd.Renderer.Send(w, r, pg)
}

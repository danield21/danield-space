package status

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var ErrNotFound = errors.New("resource not found")

//NotFoundPageHandler handles the not found page
var NotFoundPageHandler handler.Handler = handler.Chain(NotFoundHeaderHandler, NotFoundBodyLink)

var NotFoundHeaderHandler handler.Handler = view.HeaderHandler(http.StatusNotFound,
	view.Header{"Content-Type", view.HTMLContentType})

func NotFoundBodyLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := e.Repository().SiteInfo().Get(ctx)

		data := struct {
			view.BaseModel
			Message string
		}{
			view.BaseModel{
				SiteInfo: info,
			},
			"could not locate resource",
		}

		newCtx := link.PageContext(ctx, "page/status/not-found", data)

		return h(newCtx, e, w)
	}
}

func CheckNotFoundLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		var err error
		ctx, err = h(ctx, e, w)
		if err == nil {
			return ctx, nil
		} else if err != ErrNotFound {
			log.Debugf(ctx, "Error popped up: %v", err)
			return ctx, err
		}

		log.Debugf(ctx, "Error popped up: %v", err)
		return NotFoundPageHandler(ctx, e, w)
	}
}

type NotFoundHandler struct {
	Context  handler.ContextGenerator
	Renderer handler.Renderer
	SiteInfo store.SiteInfoRepository
	About    store.AboutRepository
}

func (hnd NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := hnd.Context.New(r)
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

	pg.Content = template.HTML(cnt)

	hnd.Renderer.Send(w, r, pg)
}

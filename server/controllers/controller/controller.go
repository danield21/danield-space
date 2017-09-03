package controller

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type Controller interface {
	Serve(ctx context.Context, pg *Page, rqs *http.Request) Controller
}

type ControllerHandler struct {
	Renderer Renderer
	Session  SessionGenerator
	Context  ContextGenerator
}

func (hnd ControllerHandler) ToHandler(ctr Controller) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, rqs *http.Request) {
		ctx := hnd.Context.Generate(rqs)

		pg := NewPage()
		pg.Session = hnd.Session.Generate(ctx, rqs)

		current := ctr
		for current != nil {
			current = current.Serve(ctx, pg, rqs)
		}

		pg.Session.Save(rqs, rsp)
		if pg.Status != 0 {
			rsp.WriteHeader(pg.Status)
		}
		for head, value := range pg.Header {
			rsp.Header().Add(head, value)
		}
		err := hnd.Renderer.Render(rsp, "core/page", pg)

		if err != nil {
			log.Errorf(ctx, "Error occurred during rendering %s\n%v", rqs.URL.Path, err)
		}
	})
}

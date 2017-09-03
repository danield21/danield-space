package controller

import (
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type Controller interface {
	Serve(ctx context.Context, pg *handler.Page, rqs *http.Request) Controller
}

type ControllerHandler struct {
	Renderer handler.Renderer
	Session  handler.SessionGenerator
	Context  handler.ContextGenerator
}

func (hnd ControllerHandler) ToHandler(ctr Controller) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, rqs *http.Request) {
		ctx := hnd.Context.Generate(rqs)

		pg := handler.NewPage()
		pg.Session = hnd.Session.Generate(ctx, rqs)

		current := ctr
		for current != nil {
			current = current.Serve(ctx, pg, rqs)
		}

		pg.Session.Save(rqs, rsp)
		err := hnd.Renderer.Send(rsp, rqs, pg)

		if err != nil {
			log.Errorf(ctx, "Error occurred during rendering %s\n%v", rqs.URL.Path, err)
		}
	})
}

package controller

import (
	"context"
	"encoding/json"
	"net/http"

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

		mime, err := Negotiate(rqs, "text/html", "application/json")
		if err != nil {
			log.Errorf(ctx, "Content Negotiation failed\n%v", err)
			rsp.WriteHeader(http.StatusNotAcceptable)
			rsp.Write(nil)
			return
		}

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

		if pg.Header["Content-Type"] == "" {
			rsp.Header().Add("Content-Type", mime.MIME)
		}

		switch mime.MIME {
		case "text/html":
			err := hnd.Renderer.Render(rsp, "core/page", pg)
			if err != nil {
				log.Errorf(ctx, "Error occurred during rendering %s\n%v", rqs.URL.Path, err)
			}
			break
		case "application/json":
			bPg, err := json.Marshal(pg)
			if err != nil {
				log.Errorf(ctx, "Error occurred during rendering %s\n%v", rqs.URL.Path, err)
				break
			}
			rsp.Write(bPg)
			break
		}
	})
}

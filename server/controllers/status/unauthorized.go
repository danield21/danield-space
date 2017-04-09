package status

import (
	"errors"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/view"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

var ErrUnauthorized = errors.New("unauthorized to see resource")

//NotFoundPageHandler handles the not found page
var UnauthorizedPageHandler handler.Handler = handler.Chain(UnauthorizedHeaderHandler, UnauthorizedBodyLink)

var UnauthorizedHeaderHandler handler.Handler = view.HeaderHandler(http.StatusUnauthorized,
	view.Header{"Content-Type", view.HTMLContentType},
)

func UnauthorizedBodyLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)
		r := handler.Request(ctx)

		data := struct {
			view.BaseModel
			Redirect string `json: "-"`
			Message  string
		}{
			BaseModel: view.BaseModel{
				SiteInfo: info,
			},
			Redirect: r.URL.Path,
			Message:  "unauthorized to view this resource",
		}

		newCtx := link.PageContext(ctx, "page/status/unauthorized", data)

		return h(newCtx, e, w)
	}
}

func CheckUnauthorizedLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		var err error
		ctx, err = h(ctx, e, w)
		if err == nil {
			return ctx, nil
		} else if err != ErrUnauthorized {
			return ctx, err
		}

		return UnauthorizedPageHandler(ctx, e, w)
	}
}

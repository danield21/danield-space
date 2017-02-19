package status

import (
	"errors"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"golang.org/x/net/context"
)

var ErrUnauthorized = errors.New("unauthorized to see resource")

//NotFoundPageHandler handles the not found page
var UnauthorizedPageHandler service.Handler = service.Chain(NotFoundHeaderHandler, NotFoundBodyLink)

var UnauthorizedHeaderHandler service.Handler = view.HeaderHandler(http.StatusUnauthorized,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()})

func UnauthorizedBodyLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)
		r := service.Request(ctx)

		data := struct {
			service.BaseModel
			Redirect string `json: "-"`
			Message  string
		}{
			BaseModel: service.BaseModel{
				SiteInfo: info,
			},
			Redirect: r.URL.Path,
			Message:  "unauthorized to view this resource",
		}

		newCtx := link.PageContext(ctx, "page/status/not-found", data)

		return h(newCtx, e, w)
	}
}

func CheckUnauthorizedLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
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

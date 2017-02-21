package status

import (
	"errors"
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"golang.org/x/net/context"
)

var ErrNotFound = errors.New("resource not found")

type notFoundPage struct {
	envir.Scope
	ThemeField string
	DataField  interface{}
}

func (p notFoundPage) Theme() string {
	return p.ThemeField
}

func (p notFoundPage) Data() interface{} {
	return p.DataField
}

func (p notFoundPage) Page() string {
	return "page/status/not-found"
}

//NotFoundPageHandler handles the not found page
var NotFoundPageHandler handler.Handler = handler.Chain(NotFoundHeaderHandler, NotFoundBodyLink)

var NotFoundHeaderHandler handler.Handler = view.HeaderHandler(http.StatusNotFound,
	view.Header{"Content-Type", view.HTMLContentType})

func NotFoundBodyLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		info := siteInfo.Get(ctx)

		data := struct {
			handler.BaseModel
			Message string
		}{
			handler.BaseModel{
				SiteInfo: info,
			},
			"could not locate resource",
		}

		newCtx := link.PageContext(ctx, "page/status/not-found", data)

		return h(newCtx, e, w)
	}
}

func CheckNotFoundLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		var err error
		ctx, err = h(ctx, e, w)
		if err == nil {
			return ctx, nil
		} else if err != ErrNotFound {
			return ctx, err
		}

		return NotFoundPageHandler(ctx, e, w)
	}
}

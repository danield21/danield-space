package status

import (
	"net/http"

	"errors"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
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

func NotFoundHeaderHandler(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	return view.HeaderHandler(http.StatusNotFound,
		view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()})(scp, e, w)
}

//NotFoundHandler handles the not found page
func NotFoundHandler(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusNotFound)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := service.BaseModel{
		SiteInfo: info,
	}

	return scp, e.View(w, useTheme, "page/status/not-found", pageData)
}

func NotFoundLink(h service.Handler) service.Handler {
	return func(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
		var err error
		scp, err = h(scp, e, w)
		if err == nil {
			return scp, nil
		} else if err != ErrNotFound {
			return scp, err
		}

		return NotFoundHandler(scp, e, w)
	}
}

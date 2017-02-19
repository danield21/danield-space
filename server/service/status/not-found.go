package status

import (
	"net/http"

	"errors"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
)

var ErrNotFound = errors.New("resource not found")

//NotFoundHandler handles the not found page
func NotFoundHandler(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
	r := scp.Request()
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusNotFound)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := service.BaseModel{
		SiteInfo: info,
	}

	return e.View(w, useTheme, "page/status/not-found", pageData)
}

func NotFoundLink(h service.Handler) service.Handler {
	return func(scp envir.Scope, e envir.Environment, w http.ResponseWriter) error {
		err := h(scp, e, w)
		if err == nil {
			return nil
		} else if err != ErrNotFound {
			return err
		}

		return NotFoundHandler(scp, e, w)
	}
}

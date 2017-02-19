package status

import (
	"net/http"

	"errors"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
)

var ErrUnauthorized = errors.New("unauthorized to see resource")

//Unauthorized handles the unauthorized page
func UnauthorizedHandler(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusUnauthorized)

	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := struct {
		service.BaseModel
		Redirect string
	}{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Redirect: r.URL.Path,
	}

	return scp, e.View(w, useTheme, "page/status/unauthorized", pageData)
}

func UnauthorizedLink(h service.Handler) service.Handler {
	return func(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
		var err error
		scp, err = h(scp, e, w)
		if err == nil {
			return scp, nil
		} else if err != ErrUnauthorized {
			return scp, err
		}

		return UnauthorizedHandler(scp, e, w)
	}
}

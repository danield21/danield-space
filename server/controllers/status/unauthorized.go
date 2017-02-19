package status

import (
	"errors"
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"golang.org/x/net/context"
)

var ErrUnauthorized = errors.New("unauthorized to see resource")

//Unauthorized handles the unauthorized page
func UnauthorizedHandler(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := service.Request(ctx)
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	w.WriteHeader(http.StatusUnauthorized)

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

	return ctx, e.View(w, useTheme, "page/status/unauthorized", pageData)
}

func UnauthorizedLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		var err error
		ctx, err = h(ctx, e, w)
		if err == nil {
			return ctx, nil
		} else if err != ErrUnauthorized {
			return ctx, err
		}

		return UnauthorizedHandler(ctx, e, w)
	}
}

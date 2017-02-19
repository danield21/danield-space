package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/rest"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type signinModel struct {
	service.BaseModel
	Redirect string
}

//SignInHeaders contains the headers for index
func SignInHeaders(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return ctx, nil
}

//SignIn handles the index page
func SignIn(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := service.Request(ctx)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := signinModel{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Redirect: rest.GetRedirect(r),
	}

	SignInHeaders(ctx, e, w)
	err := e.View(w, useTheme, "page/admin/signin", pageData)
	if err != nil {
		log.Errorf(ctx, "admin.SignIn - Unable to generate page:\n%v", err)
	}
	return ctx, err
}

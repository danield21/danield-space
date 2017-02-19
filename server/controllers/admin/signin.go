package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"github.com/danield21/danield-space/server/service"

	"github.com/danield21/danield-space/server/controllers/rest"
	"google.golang.org/appengine/log"
)

type signinModel struct {
	service.BaseModel
	Redirect string
}

//SignInHeaders contains the headers for index
func SignInHeaders(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	w.Header().Set("Content-Type", service.HTML.AddCharset("utf-8").String())
	return scp, nil
}

//SignIn handles the index page
func SignIn(scp envir.Scope, e envir.Environment, w http.ResponseWriter) (envir.Scope, error) {
	r := scp.Request()
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := signinModel{
		BaseModel: service.BaseModel{
			SiteInfo: info,
		},
		Redirect: rest.GetRedirect(r),
	}

	SignInHeaders(scp, e, w)
	err := e.View(w, useTheme, "page/admin/signin", pageData)
	if err != nil {
		log.Errorf(ctx, "admin.SignIn - Unable to generate page:\n%v", err)
	}
	return scp, err
}

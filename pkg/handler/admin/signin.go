package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"

	"github.com/danield21/danield-space/pkg/handler/rest"
	"google.golang.org/appengine/log"
)

type signinModel struct {
	handler.BaseModel
	Redirect string
}

//SignInHeaders contains the headers for index
func SignInHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//SignIn handles the index page
func SignIn(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := signinModel{
		BaseModel: handler.BaseModel{
			SiteInfo: info,
		},
		Redirect: rest.GetRedirect(r),
	}

	SignInHeaders(e, w, r)
	err := e.View(w, useTheme, "page/admin/signin", pageData)
	if err != nil {
		log.Errorf(ctx, "admin.SignIn - Unable to generate page:\n%v", err)
	}
}

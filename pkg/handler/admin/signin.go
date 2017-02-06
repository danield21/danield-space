package admin

import (
	"net/http"

	"github.com/danield21/danield-space/pkg/controllers/siteInfo"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler"

	"github.com/danield21/danield-space/pkg/handler/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type signinModel struct {
	SiteInfo siteInfo.SiteInfo
	Redirect string
}

//SignInHeaders contains the headers for index
func SignInHeaders(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", handler.HTML.AddCharset("utf-8").String())
}

//SignIn handles the index page
func SignIn(e envir.Environment, w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	info, _ := siteInfo.Get(ctx)

	pageData := signinModel{
		SiteInfo: info,
		Redirect: rest.GetRedirect(r),
	}

	SignInHeaders(e, w, r)
	theme := e.Theme(r)
	err := e.View(w, theme, "page/admin/signin", pageData)
	if err != nil {
		log.Errorf(ctx, "Unable to generate SignIn page:\n%v", err)
	}
}

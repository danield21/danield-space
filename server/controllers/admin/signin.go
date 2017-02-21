package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/rest"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/repository/theme"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type signinModel struct {
	handler.BaseModel
	Redirect string
}

//SignInHeaders contains the headers for index
func SignInHeaders(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Header().Set("Content-Type", view.HTMLContentType)
	return ctx, nil
}

//SignIn handles the index page
func SignIn(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
	r := handler.Request(ctx)
	useTheme := e.Theme(r, theme.GetApp(ctx))

	info := siteInfo.Get(ctx)

	pageData := signinModel{
		BaseModel: handler.BaseModel{
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

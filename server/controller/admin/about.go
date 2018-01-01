package admin

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type AboutController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	About               store.AboutRepository
	InternalServerError controller.Controller
	Unauthorized        controller.Controller
	PutAbout            Processor
}

func (ctr AboutController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	usr := user.Current(ctx)
	if usr == nil {
		log.Errorf(ctx, "%v", errors.New("user is not authenicated"))
		return ctr.Unauthorized
	}

	signOut, err := user.LogoutURL(ctx, "/")
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "cannot create a url for logging out"))
	}

	info := ctr.SiteInfo.Get(ctx)

	frm := form.NewForm()

	if rqs.Method == http.MethodPost {
		frm = ctr.PutAbout.Process(ctx, rqs, pg.Session)
	}

	if frm.IsEmpty() {
		html, err := ctr.About.Get(ctx)
		if err == nil {
			abtFld := new(form.Field)
			abtFld.Values = []string{string(html)}
			frm.Fields["about"] = abtFld
		} else {
			log.Warningf(ctx, "Unable to get about summary\n%v", err)
		}
	}

	cnt, err := ctr.Renderer.String("page/admin/about", struct {
		User    string
		Form    form.Form
		SignOut string
	}{
		User:    usr.String(),
		Form:    frm,
		SignOut: signOut,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

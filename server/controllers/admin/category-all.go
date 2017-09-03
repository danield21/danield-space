package admin

import (
	"html/template"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/controllers/controller"
	"github.com/danield21/danield-space/server/controllers/session"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
)

type CategoryAllController struct {
	Renderer            handler.Renderer
	SiteInfo            store.SiteInfoRepository
	Category            store.CategoryRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr CategoryAllController) Serve(ctx context.Context, pg *handler.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := session.User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	cats, err := ctr.Category.GetAll(ctx)
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to get all categories"))
		return ctr.InternalServerError
	}

	cnt, err := ctr.Renderer.Render(ctx, "page/admin/category-all", struct {
		User       string
		Categories []*store.Category
	}{
		User:       usr,
		Categories: cats,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

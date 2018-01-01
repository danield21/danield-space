package admin

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

type CategoryAllController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Category            store.CategoryRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr CategoryAllController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	cats, err := ctr.Category.GetAll(ctx)
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to get all categories"))
		return ctr.InternalServerError
	}

	cnt, err := ctr.Renderer.String("page/admin/category-all", struct {
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

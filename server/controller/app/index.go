package app

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type IndexController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	InternalServerError controller.Controller
}

func (ctr IndexController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	info := ctr.SiteInfo.Get(ctx)

	a, err := ctr.Article.GetAll(ctx, 10)
	if err != nil {
		log.Errorf(ctx, "app.IndexHandler - Unable to get last 10 articles\n%v", err)
	}

	cnt, err := ctr.Renderer.String("page/app/index", struct {
		Articles    []*store.Article
		Description string
	}{
		Articles:    a,
		Description: info.Description,
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

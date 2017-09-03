package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type ArticleAllController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr ArticleAllController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	arts, err := ctr.Article.GetAll(ctx, -1)
	if err != nil {
		log.Errorf(ctx, "admin.ArticleAllHandler - Unable to get all acounts\n%v", err)
	}

	cnt, err := ctr.Renderer.String("page/admin/article-all", struct {
		User     string
		Articles []*store.Article
	}{
		User:     usr,
		Articles: arts,
	})

	if err != nil {
		log.Errorf(ctx, "admin.AccountAllHandler - Unable to render content\n%v", err)
		return ctr.InternalServerError
	}

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner
	pg.Content = template.HTML(cnt)

	return nil
}

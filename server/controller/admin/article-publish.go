package admin

import (
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type ArticlePublishController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Category            store.CategoryRepository
	Unauthorized        controller.Controller
	PutArticle          Processor
	InternalServerError controller.Controller
}

func (ctr ArticlePublishController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {

	usr, signedIn := User(pg.Session)
	if !signedIn {
		return ctr.Unauthorized
	}

	info := ctr.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	frm := form.NewForm()

	if rqs.Method == http.MethodPost {
		frm = ctr.PutArticle.Process(ctx, rqs, pg.Session)
	}

	cats, err := ctr.Category.GetAll(ctx)
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to get article categories"))
		return ctr.InternalServerError
	}

	cnt, err := ctr.Renderer.String("page/admin/article-publish", struct {
		User       string
		Form       form.Form
		Categories []*store.Category
	}{
		User:       usr,
		Form:       frm,
		Categories: cats,
	})

	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "unable to render content"))
		return ctr.InternalServerError
	}

	pg.Content = template.HTML(cnt)

	return nil
}

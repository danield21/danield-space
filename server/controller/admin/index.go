package admin

import (
	"context"
	"html/template"
	"net/http"

	"github.com/danield21/danield-space/server/controller"
	"github.com/danield21/danield-space/server/store"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type IndexController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	Category            store.CategoryRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr IndexController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	usr := user.Current(ctx)
	if usr == nil {
		return ctr.Unauthorized
	}

	log.Debugf(ctx, "%v", user.IsAdmin(ctx))

	signOut, err := user.LogoutURL(ctx, "/")
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "cannot create a url for logging out"))
	}

	info := ctr.SiteInfo.Get(ctx)
	cats, _ := ctr.Category.GetAll(ctx)
	arts, _ := ctr.Article.GetAll(ctx, 1)

	cnt, err := ctr.Renderer.String("page/admin/index", struct {
		User          string
		HasCategories bool
		HasArticles   bool
		SignOut       string
	}{
		User:          usr.String(),
		HasCategories: len(cats) > 0,
		HasArticles:   len(arts) > 0,
		SignOut:       signOut,
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

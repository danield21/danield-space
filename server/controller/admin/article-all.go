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

type ArticleAllController struct {
	Renderer            controller.Renderer
	SiteInfo            store.SiteInfoRepository
	Article             store.ArticleRepository
	Unauthorized        controller.Controller
	InternalServerError controller.Controller
}

func (ctr ArticleAllController) Serve(ctx context.Context, pg *controller.Page, rqs *http.Request) controller.Controller {
	usr := user.Current(ctx)
	if usr == nil {
		return ctr.Unauthorized
	}

	signOut, err := user.LogoutURL(ctx, "/")
	if err != nil {
		log.Errorf(ctx, "%v", errors.Wrap(err, "cannot create a url for logging out"))
	}

	info := ctr.SiteInfo.Get(ctx)

	arts, err := ctr.Article.GetAll(ctx, -1)
	if err != nil {
		log.Errorf(ctx, "admin.ArticleAllHandler - Unable to get all acounts\n%v", err)
	}

	cnt, err := ctr.Renderer.String("page/admin/article-all", struct {
		User     string
		Articles []*store.Article
		SignOut  string
	}{
		User:     usr.String(),
		Articles: arts,
		SignOut:  signOut,
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

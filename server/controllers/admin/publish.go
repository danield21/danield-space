package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/view"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/gorilla/schema"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var PublishPageHandler handler.Handler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(PublishHeadersHandler),
	PublishPageLink,
	link.Theme,
)
var PublishFormHandler handler.Handler = handler.Chain(
	view.HTMLHandler,
	handler.ToLink(PublishPageHandler),
	PublishFormLink,
)

//PublishHeaders contains the headers for index
var PublishHeadersHandler handler.Handler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", view.HTMLContentType})

//Publish handles the index page
func PublishPageLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		ses := handler.Session(ctx)

		user, signedIn := GetUser(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		info := siteInfo.Get(ctx)

		cats, err := categories.GetAll(ctx)
		if err != nil {
			log.Warningf(ctx, "admin.Publish - Unable to get types of articles\n%v", err)
		}

		data := struct {
			AdminModel `json:"-"`
			Categories []*categories.Category
		}{
			AdminModel: AdminModel{
				BaseModel: handler.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Categories: cats,
		}

		return h(link.PageContext(ctx, "page/admin/publish", data), e, w)
	}
}

func PublishFormLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		ses := handler.Session(ctx)

		_, signedIn := GetUser(ses)
		if !signedIn {
			return ctx, status.ErrUnauthorized
		}

		err := r.ParseForm()
		if err != nil {
			log.Warningf(ctx, "admin.CategoryForm - Unable to parse form\n%v", err)
			return h(ctx, e, w)
		}

		var form articles.FormArticle

		decoder := schema.NewDecoder()
		err = decoder.Decode(&form, r.PostForm)
		if err != nil {
			log.Warningf(ctx, "category.Put - Unable to decode form\n%v", err)
			return h(ctx, e, w)
		}

		article, err := form.Unpack(ctx)
		if err != nil {
			log.Warningf(ctx, "category.Put - Unable unpack form\n%v", err)
			return h(ctx, e, w)
		}

		err = articles.Set(ctx, article)
		if err != nil {
			log.Warningf(ctx, "category.Put - Unable to place category into database\n%v", err)
			return h(ctx, e, w)
		}

		return h(ctx, e, w)
	}
}

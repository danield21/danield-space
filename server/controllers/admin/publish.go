package admin

import (
	"net/http"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/controllers/status"
	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/siteInfo"
	"github.com/danield21/danield-space/server/service"
	"github.com/danield21/danield-space/server/service/view"
	"github.com/gorilla/schema"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var PublishPageHandler service.Handler = service.Chain(
	view.HTMLHandler,
	service.ToLink(PublishHeadersHandler),
	PublishPageLink,
	link.Theme,
)
var PublishFormHandler service.Handler = service.Chain(
	view.HTMLHandler,
	service.ToLink(PublishPageHandler),
	PublishFormLink,
)

//PublishHeaders contains the headers for index
var PublishHeadersHandler service.Handler = view.HeaderHandler(http.StatusOK,
	view.Header{"Content-Type", service.HTML.AddCharset("utf-8").String()})

//Publish handles the index page
func PublishPageLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		ses := service.Session(ctx)

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
				BaseModel: service.BaseModel{
					SiteInfo: info,
				},
				User: user,
			},
			Categories: cats,
		}

		return h(link.PageContext(ctx, "page/admin/publish", data), e, w)
	}
}

func PublishFormLink(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) (context.Context, error) {
		r := service.Request(ctx)
		ses := service.Session(ctx)

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

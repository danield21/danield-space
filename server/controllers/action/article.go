package action

import (
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/form"
	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/articles"
	"github.com/danield21/danield-space/server/repository/categories"
	"golang.org/x/net/context"
)

const titleKey = "title"
const authorKey = "author"
const urlKey = "url"
const publishKey = "publish"
const abstractKey = "abstract"
const contentKey = "content"
const catKey = "category"

func UnpackArticle(ctx context.Context, values url.Values) (*articles.Article, form.Form) {
	var (
		err         error
		category    *categories.Category
		publishDate time.Time
		content     template.HTML
	)

	titleFld := form.NewField(titleKey, values.Get(titleKey))
	form.NotEmpty(titleFld, "title is required")

	authorFld := form.NewField(authorKey, values.Get(authorKey))
	form.NotEmpty(authorFld, "author is required")

	urlFld := form.NewField(urlKey, values.Get(urlKey))
	if form.NotEmpty(urlFld, "url is required") && !repository.ValidURLPart(urlFld.Value) {
		urlFld.ErrorMessage = "url is not in a proper format"
	}

	catFld := form.NewField(catKey, values.Get(catKey))
	if form.NotEmpty(catFld, "category is required") {
		if !repository.ValidURLPart(catFld.Value) {
			catFld.ErrorMessage = "category is not in a proper format"
		} else if category, err = categories.Get(ctx, catFld.Value); err != nil {
			catFld.ErrorMessage = "unable to find specified category"
		}
	}

	publishFld := form.NewField(publishKey, values.Get(publishKey))
	if form.NotEmpty(publishFld, "publish is required") {
		if publishDate, err = time.Parse("2006-01-02T15:04", publishFld.Value); err != nil {
			publishFld.ErrorMessage = "unable to parse time"
		}
	}

	abstractFld := form.NewField(abstractKey, values.Get(abstractKey))
	form.NotEmpty(abstractFld, "abstract is required")

	contentFld := form.NewField(contentKey, values.Get(contentKey))
	if form.NotEmpty(contentFld, "publish is required") {
		if content, err = repository.CleanHTML([]byte(contentFld.Value)); err != nil {
			contentFld.ErrorMessage = "unable to parse content"
		}
	}

	f := form.Form{
		titleFld,
		authorFld,
		catFld,
		publishFld,
		abstractFld,
		contentFld,
	}

	if f.HasErrors() {
		return nil, f
	}

	a := new(articles.Article)
	*a = articles.Article{
		Title:       titleFld.Value,
		Author:      authorFld.Value,
		Category:    category,
		URL:         urlFld.Value,
		PublishDate: publishDate,
		Abstract:    abstractFld.Value,
		HTMLContent: []byte(content),
	}

	return a, f
}

func PutArticleLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(form.WithForm(ctx, form.NewErrorForm("Unable to parse form")), e, w)
		}

		art, f := UnpackArticle(ctx, r.Form)
		if art == nil {
			return h(form.WithForm(ctx, f), e, w)
		}

		err = articles.Set(ctx, art)
		if err != nil {
			errField := form.NewField("", "")
			errField.ErrorMessage = "Unable to put into database"

			f = append(f, errField)
			return h(form.WithForm(ctx, f), e, w)
		}

		return h(form.WithForm(ctx, f), e, w)
	}
}

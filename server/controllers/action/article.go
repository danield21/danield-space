package action

import (
	"errors"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
)

const titleKey = "title"
const authorKey = "author"
const urlKey = "url"
const publishKey = "publish"
const abstractKey = "abstract"
const contentKey = "content"
const catKey = "category"

func UnpackArticle(ctx context.Context, catRepo store.CategoryRepository, values url.Values) (*store.Article, form.Form) {
	var (
		err         error
		category    *store.Category
		publishDate time.Time
		content     template.HTML
	)

	frm := form.MakeForm()
	frm.Submitted = true

	titleFld := frm.AddFieldFromValue(titleKey, values)
	form.NotEmpty(titleFld, "title is required")

	authorFld := frm.AddFieldFromValue(authorKey, values)
	form.NotEmpty(authorFld, "author is required")

	urlFld := frm.AddFieldFromValue(urlKey, values)
	if form.NotEmpty(urlFld, "url is required") && !store.ValidURLPart(urlFld.Get()) {
		form.Fail(urlFld, "url is not in a proper format")
	}

	catFld := frm.AddFieldFromValue(catKey, values)
	if form.NotEmpty(catFld, "category is required") {
		if !store.ValidURLPart(catFld.Get()) {
			form.Fail(catFld, "category is not in a proper format")
		} else if category, err = catRepo.Get(ctx, catFld.Get()); err != nil {
			form.Fail(catFld, "unable to find specified category")
		}
	}

	publishFld := frm.AddFieldFromValue(publishKey, values)
	if form.NotEmpty(publishFld, "publish is required") {
		if publishDate, err = time.Parse("2006-01-02T15:04", publishFld.Get()); err != nil {
			form.Fail(publishFld, "unable to parse time")
		}
	}

	abstractFld := frm.AddFieldFromValue(abstractKey, values)
	form.NotEmpty(abstractFld, "abstract is required")

	contentFld := frm.AddFieldFromValue(contentKey, values)
	if form.NotEmpty(contentFld, "publish is required") {
		if content, err = store.CleanHTML([]byte(contentFld.Get())); err != nil {
			form.Fail(contentFld, "unable to parse content")
		}
	}

	if frm.HasErrors() {
		return nil, frm
	}

	a := new(store.Article)
	*a = store.Article{
		Title:       titleFld.Get(),
		Author:      authorFld.Get(),
		Category:    category,
		URL:         urlFld.Get(),
		PublishDate: publishDate,
		Abstract:    abstractFld.Get(),
		HTMLContent: []byte(content),
	}

	return a, frm
}

func PutArticleLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, form.Form{Error: errors.New("Unable to parse form")}), e, w)
		}

		art, frm := UnpackArticle(ctx, e.Repository().Category(), r.Form)
		if art == nil {
			return h(WithForm(ctx, frm), e, w)
		}

		err = e.Repository().Article().Set(ctx, art)
		if err != nil {
			frm.Error = errors.New("Unable to put into database")
			return h(WithForm(ctx, frm), e, w)
		}

		return h(WithForm(ctx, frm), e, w)
	}
}

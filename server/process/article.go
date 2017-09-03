package process

import (
	"errors"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/sessions"
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

	frm := form.NewSubmittedForm()

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

type PutArticleProcessor struct {
	Article  store.ArticleRepository
	Category store.CategoryRepository
}

func (prc PutArticleProcessor) Process(ctx context.Context, req *http.Request, ses *sessions.Session) form.Form {
	err := req.ParseForm()
	if err != nil {
		return form.NewErrorForm(errors.New("Unable to parse form"))
	}

	art, frm := UnpackArticle(ctx, prc.Category, req.Form)
	if art == nil {
		return frm
	}

	err = prc.Article.Set(ctx, art)
	if err != nil {
		frm.Error = errors.New("Unable to put into database")
	}

	return frm
}

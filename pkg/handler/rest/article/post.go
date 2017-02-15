package article

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/pkg/controllers"
	"github.com/danield21/danield-space/pkg/controllers/articles"
	"github.com/danield21/danield-space/pkg/controllers/theme"
	"github.com/danield21/danield-space/pkg/envir"
	"github.com/danield21/danield-space/pkg/handler/admin"
)

type Article struct {
	controllers.DataElement
	Type        string
	Key         string
	Title       string
	PublishDate time.Time
	Abstract    string `datastore:"noindex"`
	HTMLContent []byte `datastore:"noindex"`
}

var ErrNoTitle = errors.New("No title")
var ErrNoType = errors.New("No type")
var ErrTypeBadFormat = errors.New("Bad type format")
var ErrNoKey = errors.New("No key")
var ErrKeyBadFormat = errors.New("Bad key format")
var ErrNoPublishDate = errors.New("No publish date")
var ErrNoAbstract = errors.New("No abstract")
var ErrNoContent = errors.New("No content")

func Post(e envir.Environment, w http.ResponseWriter, r *http.Request) {
	ctx := e.Context(r)
	session := e.Session(r)

	_, signed := admin.GetUser(session)
	if !signed {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	a, err := ParseForm(r)
	if err != nil {
		log.Warningf(ctx, "article.Post - Unable to parse form\n%v", err)
		return
	}

	err = articles.Set(ctx, a)
	if err != nil {
		log.Warningf(ctx, "article.Post - Unable to put article into database\n%v", err)
	}
}

func ParseForm(r *http.Request) (a articles.Article, err error) {
	err = r.ParseForm()

	if err != nil {
		return
	}

	Title := r.PostFormValue("title")
	if Title == "" {
		err = ErrNoTitle
	}

	Key, err := parseKey(r.PostForm)
	if err != nil {
		return
	}

	Type, err := parseType(r.PostForm)
	if err != nil {
		return
	}

	Publish, err := parsePublish(r.PostForm)
	if err != nil {
		return
	}

	Abstract := r.PostFormValue("abstract")
	if Abstract == "" {
		err = ErrNoAbstract
	}

	Content := r.PostFormValue("content")
	if Content == "" {
		err = ErrNoTitle
	}

	a.Title = Title
	a.Type = Type
	a.Key = Key
	a.PublishDate = Publish
	a.Abstract = Abstract
	a.HTMLContent = []byte(Content)

	return
}

func parseKey(form url.Values) (Key string, err error) {
	temp := form.Get("key")
	if temp == "" {
		err = ErrNoType
		return
	} else if !theme.ValidTheme(temp) {
		err = ErrKeyBadFormat
		return
	}
	Key = temp
	return
}

func parseType(form url.Values) (Type string, err error) {
	tempType := form.Get("type")
	if tempType == "" {
		err = ErrNoType
		return
	} else if !theme.ValidTheme(tempType) {
		err = ErrTypeBadFormat
		return
	}
	Type = tempType
	return
}

func parsePublish(form url.Values) (Time time.Time, err error) {
	temp := form.Get("publish")
	if temp == "" {
		err = ErrNoPublishDate
		return
	}

	Time, err = time.Parse("2006-01-02T15:04", temp)
	return
}

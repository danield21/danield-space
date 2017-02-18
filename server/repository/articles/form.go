package articles

import (
	"errors"
	"time"

	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/categories"
	"github.com/danield21/danield-space/server/repository/theme"
	"golang.org/x/net/context"
)

//FormArticle contains information about articles written on this website.
type FormArticle struct {
	Title       string `schema:"title"`
	Author      string `schema:"author"`
	Url         string `schema:"url"`
	PublishDate string `schema:"publish"`
	Abstract    string `schema:"abstract"`
	Content     string `schema:"content"`
	Category    string `schema:"category"`
}

var ErrNoTitle = errors.New("No title")
var ErrCategoryBadFormat = errors.New("Bad category format")
var ErrUrlBadFormat = errors.New("Bad url format")
var ErrNoPublishDate = errors.New("No publish date")
var ErrNoAbstract = errors.New("No abstract")
var ErrNoContent = errors.New("No content")

func (f FormArticle) Unpack(ctx context.Context) (*Article, error) {

	if !theme.ValidTheme(f.Url) {
		return nil, ErrUrlBadFormat
	}

	category, err := parseCategory(ctx, f.Category)
	if err != nil {
		return nil, err
	}

	publish, err := parsePublish(f.PublishDate)
	if err != nil {
		return nil, err
	}

	content, err := parseContent(f.Content)
	if err != nil {
		return nil, err
	}

	a := new(Article)
	a.Title = f.Title
	a.Author = f.Author
	a.Url = f.Url
	a.PublishDate = publish
	a.Abstract = f.Abstract
	a.HTMLContent = content
	a.Category = category

	return a, nil
}

func parseUrl(url string) (string, error) {
	if !theme.ValidTheme(url) {
		return "", ErrUrlBadFormat
	}
	return url, nil
}

func parseCategory(ctx context.Context, catUrl string) (*categories.Category, error) {
	if !theme.ValidTheme(catUrl) {
		return nil, ErrCategoryBadFormat
	}
	return categories.Get(ctx, catUrl)
}

func parsePublish(publish string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", publish)
}

func parseContent(content string) ([]byte, error) {
	html, err := repository.CleanHTML([]byte(content))
	return []byte(html), err
}

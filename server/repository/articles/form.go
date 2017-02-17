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
	Title       string
	Url         string
	PublishDate string
	Abstract    string
	Content     string
	Category    string
}

var ErrNoTitle = errors.New("No title")
var ErrCategoryBadFormat = errors.New("Bad category format")
var ErrUrlBadFormat = errors.New("Bad url format")
var ErrNoPublishDate = errors.New("No publish date")
var ErrNoAbstract = errors.New("No abstract")
var ErrNoContent = errors.New("No content")

func (f FormArticle) Unpack(ctx context.Context) (a Article, err error) {
	var (
		category categories.Category
		publish  time.Time
		content  []byte
	)

	if err != nil {
		return
	}

	if !theme.ValidTheme(f.Url) {
		err = ErrUrlBadFormat
		return
	}

	category, err = parseCategory(ctx, f.Category)
	if err != nil {
		return
	}

	publish, err = parsePublish(f.PublishDate)
	if err != nil {
		return
	}

	content, err = parseContent(f.Content)
	if err != nil {
		return
	}

	a = Article{
		Title:       f.Title,
		Url:         f.Url,
		PublishDate: publish,
		Abstract:    f.Abstract,
		HTMLContent: content,
		Category:    category,
	}

	return
}

func parseUrl(url string) (cleanUrl string, err error) {
	if !theme.ValidTheme(url) {
		err = ErrUrlBadFormat
		return
	}
	cleanUrl = url
	return
}

func parseCategory(ctx context.Context, catUrl string) (category categories.Category, err error) {
	if !theme.ValidTheme(catUrl) {
		err = ErrCategoryBadFormat
		return
	}
	category, _, err = categories.Get(ctx, catUrl)
	return
}

func parsePublish(publish string) (publishDate time.Time, err error) {
	publishDate, err = time.Parse("2006-01-02T15:04", publish)
	return
}

func parseContent(content string) ([]byte, error) {
	html, err := repository.CleanHTML([]byte(content))
	return []byte(html), err
}

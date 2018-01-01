package store

import (
	"html/template"
	"log"
	"strings"
	"time"

	"context"
)

//Article contains information about articles written on this website.
type Article struct {
	DataElement
	Title       string
	Author      string
	URL         string
	PublishDate time.Time
	Abstract    string    `datastore:",noindex"`
	HTMLContent []byte    `datastore:",noindex"`
	Category    *Category `datastore:"-"`
}

//Path returns the path for a article.
func (a Article) Path() string {
	return "/" + a.Category.URL + "/" + a.URL
}

//Heading returns a heading for the article using the Type and Title.
func (a Article) Heading() string {
	return strings.Title(a.Category.Title + ": " + a.Title)
}

//Content returns HTMLContent after it has been parsed and cleaned
func (a Article) Content() (content template.HTML) {
	var err error

	content, err = CleanHTML(a.HTMLContent)
	if err != nil {
		log.Printf("WARNING: article.Content - Unable to clean HTML\n%v", err)
	}
	return
}

type ArticleRepository interface {
	Get(ctx context.Context, cat *Category, url string) (*Article, error)
	GetAll(ctx context.Context, limit int) ([]*Article, error)
	GetAllByCategory(ctx context.Context, cat *Category, limit int) ([]*Article, error)
	GetMapKeyedByCategory(ctx context.Context, Limit int) (map[*Category][]*Article, error)
	Set(ctx context.Context, article *Article) error
}

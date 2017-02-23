package articles

import (
	"log"
	"strings"
	"time"

	"html/template"

	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/categories"
)

//Article contains information about articles written on this website.
type Article struct {
	repository.DataElement
	Title       string
	Author      string
	URL         string
	PublishDate time.Time
	Abstract    string               `datastore:",noindex"`
	HTMLContent []byte               `datastore:",noindex"`
	Category    *categories.Category `datastore:"-"`
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

	content, err = repository.CleanHTML(a.HTMLContent)
	if err != nil {
		log.Printf("WARNING: article.Content - Unable to clean HTML\n%v", err)
	}
	return
}

package articles

import (
	"log"
	"strings"
	"time"

	"html/template"

	"github.com/danield21/danield-space/pkg/controllers"
	"github.com/danield21/danield-space/pkg/controllers/categories"
)

//Article contains information about articles written on this website.
type Article struct {
	controllers.DataElement
	Url         string
	Title       string
	PublishDate time.Time
	Abstract    string              `datastore:",noindex"`
	HTMLContent []byte              `datastore:",noindex"`
	Category    categories.Category `datastore:"-"`
}

//Path returns the path for a article.
func (a Article) Path() string {
	return "/" + a.Category.Url + "/" + a.Url
}

//Heading returns a heading for the article using the Type and Title.
func (a Article) Heading() string {
	return strings.Title(a.Category.Title + ": " + a.Title)
}

//Content returns HTMLContent after it has been parsed and cleaned
func (a Article) Content() (content template.HTML) {
	var err error

	content, err = controllers.CleanHTML(a.HTMLContent)
	if err != nil {
		log.Printf("WARNING: article.Content - Unable to clean HTML\n%v", err)
	}
	return
}

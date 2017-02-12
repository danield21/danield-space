package articles

import (
	"log"
	"strings"

	"html/template"

	"github.com/danield21/danield-space/pkg/controllers"
)

//Article contains information about articles written on this website.
type Article struct {
	controllers.DataElement
	Type        string
	Key         string
	Title       string
	Abstract    string `datastore:"noindex"`
	HTMLContent []byte `datastore:"noindex"`
}

//Path returns the path for a article.
func (a Article) Path() string {
	return "/" + a.Type + "/" + a.Key
}

//Heading returns a heading for the article using the Type and Title.
func (a Article) Heading() string {
	return strings.Title(a.Type + ": " + a.Title)
}

//Content returns HTMLContent after it has been parsed and cleaned
func (a Article) Content() (content template.HTML) {
	var err error

	content, err = controllers.CleanHTML(a.HTMLContent)
	if err != nil {
		log.Printf("article.Content - Unable to clean HTML\n%v", err)
	}
	return
}

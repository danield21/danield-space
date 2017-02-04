package articles

import (
	"strings"

	"html/template"

	"github.com/danield21/danield-space/pkg/controllers"
)

//Article contains information about articles written on this website.
type Article struct {
	controllers.DataElement
	Type            string
	Key             string
	Title           string
	Abstract        string
	MarkdownContent string
	HTMLContent     template.HTML
}

//Path returns the path for a article.
func (a Article) Path() string {
	return "/" + a.Type + "/" + a.Key
}

//Heading returns a heading for the article using the Type and Title.
func (a Article) Heading() string {
	return strings.Title(a.Type + ": " + a.Title)
}

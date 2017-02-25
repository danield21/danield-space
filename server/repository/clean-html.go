package repository

import (
	"bytes"
	"html/template"

	"regexp"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

var languageRegexp *regexp.Regexp = regexp.MustCompile("langauge-[\\w\\-]+")

func CleanHTML(dirtyHTML []byte) (cleanHTML template.HTML, err error) {

	reader := bytes.NewReader(dirtyHTML)
	htmlNodes, pErr := html.Parse(reader)
	if pErr != nil {
		err = pErr
		return
	}

	var renderedHTML bytes.Buffer
	html.Render(&renderedHTML, htmlNodes)

	policy := bluemonday.NewPolicy()
	policy.AllowElements(
		"i", "b", "strong", "em",
		"a", "p",
		"h1", "h2", "h3", "h4", "h5", "h6",
		"pre", "code",
		"li", "ol", "ul",
	)
	policy.AllowAttrs("href").OnElements("a")
	policy.AllowAttrs("class").Matching(languageRegexp).OnElements("code")
	policy.RequireParseableURLs(true)
	policy.AllowRelativeURLs(true)
	policy.RequireNoFollowOnFullyQualifiedLinks(true)
	policy.SanitizeBytes(renderedHTML.Bytes())
	cleanHTML = template.HTML(dirtyHTML)
	return
}

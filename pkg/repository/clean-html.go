package repository

import (
	"bytes"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

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
	policy.AllowElements("i", "b", "strong", "em", "a", "p")
	policy.AllowAttrs("href").OnElements("a")
	policy.RequireParseableURLs(true)
	policy.AllowRelativeURLs(true)
	policy.RequireNoFollowOnFullyQualifiedLinks(true)
	cleanBytes := policy.SanitizeBytes(renderedHTML.Bytes())
	cleanHTML = template.HTML(cleanBytes)
	return
}

package action

import "html/template"

type URL struct {
	URL   string
	Title string
}

func (u URL) HTML() template.HTML {
	return template.HTML("<a href=\"" + u.URL + "\">" + u.Title + "</a>")
}

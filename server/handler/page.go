package handler

import (
	"html/template"

	"golang.org/x/text/language"
)

type Page struct {
	Status   int
	Header   map[string]string
	Title    string
	Language language.Tag
	Meta     map[string]string
	Content  template.HTML
}

func NewPage() *Page {
	p := new(Page)
	p.Header = make(map[string]string)
	p.Meta = make(map[string]string)
	p.Language = language.AmericanEnglish
	return p
}

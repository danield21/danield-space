package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
	"golang.org/x/text/language"
)

type Page struct {
	Status   int
	Header   map[string]string
	Session  *sessions.Session
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

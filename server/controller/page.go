package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
	"golang.org/x/text/language"
)

type Page struct {
	Status   int               `json:"-"`
	Header   map[string]string `json:"-"`
	Session  *sessions.Session `json:"-"`
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

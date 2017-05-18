package models

import (
	"html/template"
	"log"

	"golang.org/x/net/context"
)

var DefaultAbout = []byte(
	"<section>My name is Daniel Juan Dominguez, and I am a developer.</section>",
)

type About []byte

func (a About) Content() template.HTML {
	content, err := CleanHTML(a)
	if err != nil {
		log.Printf("WARNING: About.Content - Unable to clean HTML\n%v", err)
	}
	return content
}

type AboutRepository interface {
	Get(ctx context.Context) (template.HTML, error)
	Set(ctx context.Context, abt []byte) error
}

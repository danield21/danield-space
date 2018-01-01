package store

import (
	"html/template"
	"log"

	"context"
)

var DefaultAbout = []byte(
	"<div>My name is Daniel Juan Dominguez, and I am a developer.</div>",
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

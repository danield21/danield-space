package models

import (
	"html/template"

	"golang.org/x/net/context"
)

var DefaultAbout = []byte(
	"<section>My name is Daniel Juan Dominguez, and I am a developer.</section>",
)

type AboutRepository interface {
	Get(ctx context.Context) (template.HTML, error)
	Set(ctx context.Context, abt []byte) error
}

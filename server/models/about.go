package models

import (
	"html/template"

	"golang.org/x/net/context"
)

type AboutRepository interface {
	Get(ctx context.Context) (template.HTML, error)
	Set(ctx context.Context, abt []byte) error
}

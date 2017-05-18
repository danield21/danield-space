package store

import (
	"golang.org/x/net/context"
)

type Category struct {
	DataElement
	Title       string
	URL         string
	Description string
}

func NewEmptyCategory(url string) *Category {
	cat := new(Category)
	cat.URL = url
	return cat
}

type CategoryRepository interface {
	Get(c context.Context, url string) (*Category, error)
	GetAll(c context.Context) ([]*Category, error)
	Set(ctx context.Context, cat *Category) error
	Remove(c context.Context, cat *Category) error
	IsUnique(c context.Context, category *Category) bool
}

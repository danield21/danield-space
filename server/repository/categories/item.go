package categories

import (
	"github.com/danield21/danield-space/server/repository"
)

type Category struct {
	repository.DataElement
	Title       string
	URL         string
	Description string
}

func NewEmptyCategory(url string) *Category {
	cat := new(Category)
	cat.URL = url
	return cat
}

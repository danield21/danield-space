package categories

import (
	"github.com/danield21/danield-space/server/repository"
)

type Category struct {
	repository.DataElement
	Title       string
	Url         string
	Description string
}

func EmptyCategory(url string) *Category {
	cat := new(Category)
	cat.Url = url
	return cat
}

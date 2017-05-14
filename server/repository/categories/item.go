package categories

import (
	"github.com/danield21/danield-space/server/models"
)

type Category struct {
	models.DataElement
	Title       string
	URL         string
	Description string
}

func NewEmptyCategory(url string) *Category {
	cat := new(Category)
	cat.URL = url
	return cat
}

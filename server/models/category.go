package models

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

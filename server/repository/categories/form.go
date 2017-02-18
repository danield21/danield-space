package categories

import (
	"errors"

	"github.com/danield21/danield-space/server/repository"
)

type FormCategory struct {
	Title       string `schema:"title"`
	Url         string `schema:"url"`
	Description string `schema:"description"`
}

var ErrInvalidUrl = errors.New("Url is not in a proper format")

func (f FormCategory) Unpack() (*Category, error) {
	if !repository.ValidUrlPart(f.Url) {
		return nil, ErrInvalidUrl
	}

	category := new(Category)
	category.Title = f.Title
	category.Url = f.Url
	category.Description = f.Description
	return category, nil
}

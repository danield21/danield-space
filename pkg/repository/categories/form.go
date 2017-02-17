package categories

import (
	"errors"

	"github.com/danield21/danield-space/pkg/repository"
)

type FormCategory struct {
	Title       string `schema:"title"`
	Url         string `schema:"url"`
	Description string `schema:"description"`
}

var ErrInvalidUrl = errors.New("Url is not in a proper format")

func (f FormCategory) Unpack() (category Category, err error) {
	if !repository.ValidUrlPart(f.Url) {
		err = ErrInvalidUrl
		return
	}

	category = Category{
		Title: f.Title,
		Url:   f.Url,
	}
	return
}

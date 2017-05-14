package about

import (
	"html/template"

	"github.com/danield21/danield-space/server/models"
	"github.com/danield21/danield-space/server/repository"
	"golang.org/x/net/context"
)

var DefaultAbout = []byte(
	"<section>My name is Daniel Juan Dominguez, and I am a developer.</section>",
)

const bucketKey = "about-page"

var bucket = repository.Bucket{}

func Get(c context.Context) (template.HTML, error) {
	item, err := aboutToItem(DefaultAbout)
	if err != nil {
		return "", err
	}

	field := bucket.Default(c, item)

	return itemToAbout(field)
}

func Set(c context.Context, about []byte) error {
	item, err := aboutToItem(about)
	if err != nil {
		return err
	}
	err = bucket.Set(c, item)
	return err
}

func aboutToItem(about []byte) (*models.Item, error) {
	clean, err := repository.CleanHTML(about)
	return models.NewItem(bucketKey, string(clean), "[]byte"), err
}

func itemToAbout(item *models.Item) (template.HTML, error) {
	return repository.CleanHTML([]byte(item.Value))
}

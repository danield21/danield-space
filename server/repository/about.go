package repository

import (
	"html/template"

	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
)

var DefaultAbout = []byte(
	"<section>My name is Daniel Juan Dominguez, and I am a developer.</section>",
)

const bucketKey = "about-page"

type About struct {
	Bucket models.BucketRepository
}

func (a About) Get(c context.Context) (template.HTML, error) {
	item, err := aboutToItem(DefaultAbout)
	if err != nil {
		return "", err
	}

	field := a.Bucket.Default(c, item)

	return itemToAbout(field)
}

func (a About) Set(c context.Context, about []byte) error {
	item, err := aboutToItem(about)
	if err != nil {
		return err
	}
	err = a.Bucket.Set(c, item)
	return err
}

func aboutToItem(about []byte) (*models.Item, error) {
	clean, err := CleanHTML(about)
	return models.NewItem(bucketKey, string(clean), "[]byte"), err
}

func itemToAbout(item *models.Item) (template.HTML, error) {
	return CleanHTML([]byte(item.Value))
}
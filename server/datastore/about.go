package datastore

import (
	"html/template"

	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
)

const bucketKey = "about-page"

type About struct {
	Bucket models.BucketRepository
}

func (a About) Get(c context.Context) (template.HTML, error) {
	item, err := aboutToItem(models.DefaultAbout)
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
	clean, err := models.CleanHTML(about)
	return models.NewItem(bucketKey, string(clean), "[]byte"), err
}

func itemToAbout(item *models.Item) (template.HTML, error) {
	return models.CleanHTML([]byte(item.Value))
}

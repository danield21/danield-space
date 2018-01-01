package datastore

import (
	"context"
	"html/template"

	"github.com/danield21/danield-space/server/store"
)

const bucketKey = "about-page"

type About struct {
	Bucket store.BucketRepository
}

func (a About) Get(c context.Context) (template.HTML, error) {
	item, err := aboutToItem(store.DefaultAbout)
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

func aboutToItem(about []byte) (*store.Item, error) {
	clean, err := store.CleanHTML(about)
	return store.NewItem(bucketKey, string(clean), "[]byte"), err
}

func itemToAbout(item *store.Item) (template.HTML, error) {
	return store.CleanHTML([]byte(item.Value))
}

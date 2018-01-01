package datastore

import (
	"context"
	"errors"

	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/datastore"
)

const categoryEntity = "Categories"

var ErrNoCategory = errors.New("no category found")
var ErrNilCategory = errors.New("category was nil")

type Category struct{}

func (Category) Get(ctx context.Context, url string) (*store.Category, error) {
	var categories []*store.Category

	q := datastore.NewQuery(categoryEntity).Filter("URL =", url).Limit(1)

	keys, err := q.GetAll(ctx, &categories)
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, ErrNoCategory
	}

	categories[0].Key = keys[0]
	return categories[0], nil
}

func (Category) GetAll(ctx context.Context) ([]*store.Category, error) {
	var categories []*store.Category
	q := datastore.NewQuery(categoryEntity)
	keys, err := q.GetAll(ctx, &categories)

	if err != nil {
		return nil, err
	}

	for i := range categories {
		categories[i].Key = keys[i]
	}

	return categories, nil
}

func (c Category) Set(ctx context.Context, cat *store.Category) error {
	if cat == nil {
		return ErrNilCategory
	}
	oldCat, err := c.Get(ctx, cat.URL)

	if err != nil {
		cat.DataElement = store.WithNew(store.WithPerson(ctx))
		cat.Key = datastore.NewIncompleteKey(ctx, categoryEntity, nil)
	} else {
		cat.DataElement = store.WithOld(store.WithPerson(ctx), oldCat.DataElement)
	}

	cat.Key, err = datastore.Put(ctx, cat.Key, cat)

	return err
}

func (c Category) Remove(ctx context.Context, cat *store.Category) error {
	var err error
	if cat == nil {
		return ErrNilCategory
	} else if cat.Key == nil {
		cat, err = c.Get(ctx, cat.URL)
		if err != nil {
			return err
		}
	}

	return datastore.Delete(ctx, cat.Key)
}

func (c Category) IsUnique(ctx context.Context, category *store.Category) bool {
	_, err := c.Get(ctx, category.URL)
	return err == nil
}

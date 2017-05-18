package repository

import (
	"errors"

	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const categoryEntity = "Categories"

var ErrNoCategory = errors.New("no category found")
var ErrNilCategory = errors.New("category was nil")

type Category struct{}

func (Category) Get(ctx context.Context, url string) (*models.Category, error) {
	var categories []*models.Category

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

func (Category) GetAll(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category
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

func (c Category) Set(ctx context.Context, cat *models.Category) error {
	if cat == nil {
		return ErrNilCategory
	}
	oldCat, err := c.Get(ctx, cat.URL)

	if err != nil {
		cat.DataElement = models.WithNew(models.WithPerson(ctx))
		cat.Key = datastore.NewIncompleteKey(ctx, categoryEntity, nil)
	} else {
		cat.DataElement = models.WithOld(models.WithPerson(ctx), oldCat.DataElement)
	}

	cat.Key, err = datastore.Put(ctx, cat.Key, cat)

	return err
}

func (c Category) Remove(ctx context.Context, cat *models.Category) error {
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

func (c Category) IsUnique(ctx context.Context, category *models.Category) bool {
	_, err := c.Get(ctx, category.URL)
	return err == nil
}

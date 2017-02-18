package categories

import (
	"errors"

	"github.com/danield21/danield-space/server/repository"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const entity = "Categories"

var ErrNoMatch = errors.New("no category found")
var ErrNilCategory = errors.New("category was nil")

func Get(c context.Context, url string) (*Category, error) {
	var categories []*Category

	q := datastore.NewQuery(entity).Filter("Url =", url).Limit(1)

	keys, err := q.GetAll(c, &categories)
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, ErrNoMatch
	}

	categories[0].Key = keys[0]
	return categories[0], nil
}

func GetAll(c context.Context) ([]*Category, error) {
	var categories []*Category
	q := datastore.NewQuery(entity)
	keys, err := q.GetAll(c, &categories)

	if err != nil {
		return nil, err
	}

	for i := range categories {
		categories[i].Key = keys[i]
	}

	return categories, nil
}

func Set(ctx context.Context, cat *Category) error {
	if cat == nil {
		return ErrNilCategory
	}
	oldCat, err := Get(ctx, cat.Url)

	if err != nil {
		cat.Key = datastore.NewIncompleteKey(ctx, entity, nil)
		cat.DataElement = repository.WithNew(repository.WithPerson(ctx))
	} else {
		cat.DataElement = repository.WithOld(repository.WithPerson(ctx), oldCat.DataElement)
	}

	cat.Key, err = datastore.Put(ctx, cat.Key, &cat)

	return err
}

func Remove(c context.Context, cat *Category) error {
	var err error
	if cat == nil {
		return ErrNilCategory
	} else if cat.Key == nil {
		cat, err = Get(c, cat.Url)
		if err != nil {
			return err
		}
	}

	return datastore.Delete(c, cat.Key)
}

func IsUnique(c context.Context, category *Category) bool {
	_, err := Get(c, category.Url)
	return err == nil
}

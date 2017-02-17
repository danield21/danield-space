package categories

import (
	"errors"

	"github.com/danield21/danield-space/pkg/controllers"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const entity = "Categories"

var ErrNoMatch = errors.New("No category found")

func Get(c context.Context, url string) (category Category, key *datastore.Key, err error) {
	var (
		categories []Category
		keys       []*datastore.Key
	)

	q := datastore.NewQuery(entity).Filter("Url =", url).Limit(1)
	keys, err = q.GetAll(c, &categories)

	if err != nil {
		return
	}

	if len(keys) == 0 {
		err = ErrNoMatch
		return
	}

	category = categories[0]
	key = keys[0]
	return
}

func GetAll(c context.Context) (categories []Category, err error) {
	q := datastore.NewQuery(entity)
	_, err = q.GetAll(c, &categories)

	return
}

func Set(c context.Context, category Category) (err error) {
	oldCategory, key, dErr := Get(c, category.Url)

	if dErr != nil {
		key = datastore.NewIncompleteKey(c, entity, nil)
		category.DataElement = controllers.WithNew("site")
	} else {
		category.DataElement = controllers.WithOld(oldCategory.DataElement, "site")
	}

	_, err = datastore.Put(c, key, &category)

	return
}

func Remove(c context.Context, category Category) (err error) {
	_, key, dErr := Get(c, category.Url)

	if dErr != nil {
		return
	}

	err = datastore.Delete(c, key)

	return
}

func IsUnique(c context.Context, category Category) bool {
	_, _, err := Get(c, category.Url)
	return err == nil
}

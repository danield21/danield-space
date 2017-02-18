package bucket

import (
	"github.com/danield21/danield-space/server/repository"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const entity = "Bucket"

//Get gets an item from the bucket with the same field
func Get(ctx context.Context, field string) (*Item, error) {
	var items []*Item

	q := datastore.NewQuery(entity).Filter("Field =", field).Limit(1)

	keys, err := q.GetAll(ctx, &items)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrFieldNotFound
	}

	items[0].Key = keys[0]
	return items[0], nil
}

//GetAll gets all items with the fields listed.
//If there are fields missing, then
func GetAll(ctx context.Context, fields ...string) (have []*Item, missing []string) {

	for _, f := range fields {
		item, err := Get(ctx, f)
		if err != nil {
			missing = append(missing, f)
		} else {
			have = append(have, item)
		}
	}

	return
}

//Set sets the field with item
func Set(ctx context.Context, item *Item) error {
	if item == nil {
		return ErrNilItem
	}

	oldItem, err := Get(ctx, item.Field)

	if err != nil {
		log.Warningf(ctx, "bucket.Set - Unable to get previous item, creating new\n%v", err)

		item.DataElement = repository.WithNew(repository.WithPerson(ctx))
		item.Key = datastore.NewIncompleteKey(ctx, entity, nil)
	} else {
		item.DataElement = repository.WithOld(repository.WithPerson(ctx), oldItem.DataElement)
	}

	item.Key, err = datastore.Put(ctx, item.Key, item)
	return err
}

//SetAll sets all items
//No logic is put in place for items with duplicate field,
//so field will be overwritten as it loops through.
//Latest value survives
//No transaction, so if any fail, it will not rollback any successful
func SetAll(ctx context.Context, items ...*Item) error {
	var (
		fields []string
		keys   []*datastore.Key
		err    error
	)

	if items == nil {
		return nil
	}

	for _, item := range items {
		if item == nil {
			return nil
		}
		fields = append(fields, item.Field)
	}

	have, missing := GetAll(ctx, fields...)

CheckingForNew:
	for _, i := range items {
		for _, m := range missing {
			if m != i.Field {
				continue
			}

			i.DataElement = repository.WithNew(repository.WithPerson(ctx))
			i.Key = datastore.NewIncompleteKey(ctx, entity, nil)
			have = append(have, i)

			continue CheckingForNew
		}

		for _, h := range have {
			if h.Field != i.Field {
				continue
			}

			i.DataElement = repository.WithOld(repository.WithPerson(ctx), h.DataElement)
		}
	}

	for _, i := range have {
		keys = append(keys, i.Key)
	}

	keys, err = datastore.PutMulti(ctx, keys, have)
	if err != nil {
		return err
	}

	for i, k := range keys {
		have[i].Key = k
	}

	return nil
}

func Default(ctx context.Context, defaultItem *Item) *Item {
	var items []*Item

	if defaultItem == nil {
		return nil
	}

	q := datastore.NewQuery(entity).Filter("Field =", defaultItem.Field).Limit(1)
	_, err := q.GetAll(ctx, &items)
	if err != nil || len(items) > 0 {
		return items[0]
	}

	log.Infof(ctx, "Field %s missing, using default %s", defaultItem.Field, defaultItem.Value)
	key := datastore.NewIncompleteKey(ctx, entity, nil)
	defaultItem.DataElement = repository.WithNew("site")

	if key, err := datastore.Put(ctx, key, defaultItem); err != nil {
		defaultItem.Key = key
	}

	return defaultItem
}

func DefaultAll(ctx context.Context, defaultItems ...*Item) []*Item {
	if defaultItems == nil {
		return nil
	}

	var (
		fields       []string
		missingItems []*Item
		keys         []*datastore.Key
		err          error
	)

	for _, item := range defaultItems {
		if item == nil {
			return nil
		}
		fields = append(fields, item.Field)
	}

	have, missing := GetAll(ctx, fields...)

CheckingForNew:
	for _, m := range missing {
		for _, i := range defaultItems {
			if m != i.Field {
				continue
			}
			log.Infof(ctx, "Field %s missing, using default %s", m, i.Value)

			var newItem *Item
			*newItem = *i

			newItem.DataElement = repository.WithNew("site")
			missingItems = append(missingItems, newItem)
			have = append(have, newItem)
			keys = append(keys, datastore.NewIncompleteKey(ctx, entity, nil))

			continue CheckingForNew
		}
	}

	keys, err = datastore.PutMulti(ctx, keys, missingItems)
	if err != nil {
		log.Infof(ctx, "bucket.DefaultAll - Unable to put missing info into database\n%v", err)
		return have
	}

	for i, k := range keys {
		missingItems[i].Key = k
	}

	return have
}

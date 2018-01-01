package datastore

import (
	"context"
	"errors"

	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Bucket struct{}

const bucketEntity = "Bucket"

//ErrFieldNotFound appears when the field request does not exist in the database
var ErrFieldNotFound = errors.New("field not found")

//ErrNilItem appears when item parameter is nil
var ErrNilItem = errors.New("err was nil")

//Get gets an item from the bucket with the same field
func (b Bucket) Get(ctx context.Context, field string) (*store.Item, error) {
	var items []*store.Item

	q := datastore.NewQuery(bucketEntity).Filter("Field =", field).Limit(1)

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
func (b Bucket) GetAll(ctx context.Context, fields ...string) (have []*store.Item, missing []string) {

	for _, f := range fields {
		item, err := b.Get(ctx, f)
		if err != nil {
			missing = append(missing, f)
		} else {
			have = append(have, item)
		}
	}

	return
}

//Set sets the field with item
func (b Bucket) Set(ctx context.Context, item *store.Item) error {
	if item == nil {
		return ErrNilItem
	}

	oldItem, err := b.Get(ctx, item.Field)

	if err != nil {
		log.Warningf(ctx, "bucket.Set - Unable to get previous item, creating new\n%v", err)

		item.DataElement = store.WithNew(store.WithPerson(ctx))
		item.Key = datastore.NewIncompleteKey(ctx, bucketEntity, nil)
	} else {
		item.DataElement = store.WithOld(store.WithPerson(ctx), oldItem.DataElement)
	}

	item.Key, err = datastore.Put(ctx, item.Key, item)
	return err
}

//SetAll sets all items
//No logic is put in place for items with duplicate field,
//so field will be overwritten as it loops through.
//Latest value survives
//No transaction, so if any fail, it will not rollback any successful
func (b Bucket) SetAll(ctx context.Context, items ...*store.Item) error {
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

	have, missing := b.GetAll(ctx, fields...)

CheckingForNew:
	for _, i := range items {
		for _, m := range missing {
			if m != i.Field {
				continue
			}

			i.DataElement = store.WithNew(store.WithPerson(ctx))
			i.Key = datastore.NewIncompleteKey(ctx, bucketEntity, nil)
			have = append(have, i)

			continue CheckingForNew
		}

		for _, h := range have {
			if h.Field != i.Field {
				continue
			}

			i.DataElement = store.WithOld(store.WithPerson(ctx), h.DataElement)
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

func (b Bucket) Default(ctx context.Context, defaultItem *store.Item) *store.Item {
	var items []*store.Item

	if defaultItem == nil {
		return nil
	}

	q := datastore.NewQuery(bucketEntity).Filter("Field =", defaultItem.Field).Limit(1)
	_, err := q.GetAll(ctx, &items)
	if err != nil || len(items) > 0 {
		return items[0]
	}

	log.Infof(ctx, "Field %s missing, using default %s", defaultItem.Field, defaultItem.Value)
	key := datastore.NewIncompleteKey(ctx, bucketEntity, nil)
	defaultItem.DataElement = store.WithNew("site")

	if key, err := datastore.Put(ctx, key, defaultItem); err != nil {
		defaultItem.Key = key
	}

	return defaultItem
}

func (b Bucket) DefaultAll(ctx context.Context, defaultItems ...*store.Item) []*store.Item {
	if defaultItems == nil {
		return nil
	}

	var (
		fields       []string
		missingItems []*store.Item
		keys         []*datastore.Key
		err          error
	)

	for _, item := range defaultItems {
		if item == nil {
			return nil
		}
		fields = append(fields, item.Field)
	}

	have, missing := b.GetAll(ctx, fields...)

CheckingForNew:
	for _, m := range missing {
		for _, i := range defaultItems {
			if m != i.Field {
				continue
			}
			log.Infof(ctx, "Field %s missing, using default \"%s\"", m, i.Value)

			newItem := new(store.Item)
			*newItem = *i

			newItem.DataElement = store.WithNew("site")
			missingItems = append(missingItems, newItem)
			have = append(have, newItem)
			keys = append(keys, datastore.NewIncompleteKey(ctx, bucketEntity, nil))

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

package bucket

import (
	"fmt"

	"github.com/danield21/danield-space/pkg/repository"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const entity = "Bucket"

//Get gets an item from the bucket with the same field
func Get(c context.Context, field string) (item Item, key *datastore.Key, err error) {
	var (
		items []Item
		keys  []*datastore.Key
	)

	q := datastore.NewQuery(entity).Filter("Field =", field).Limit(1)

	keys, err = q.GetAll(c, &items)
	if err != nil {
		return
	}

	if len(items) == 0 {
		err = ErrFieldNotFound
	} else {
		item = items[0]
		key = keys[0]
	}

	return
}

//GetAll gets all items with the fields listed
func GetAll(c context.Context, fields ...string) (items []Item, keys []*datastore.Key, err error) {
	for _, f := range fields {
		item, key, dErr := Get(c, f)
		if dErr != nil {
			if err != nil {
				err = fmt.Errorf("%v\n%v", err, dErr)
			} else {
				err = dErr
			}

			continue
		}
		items = append(items, item)
		keys = append(keys, key)
	}

	return
}

//Set sets the field with item
func Set(c context.Context, item Item) (err error) {
	oldItem, key, dErr := Get(c, item.Field)

	if dErr != nil {
		log.Warningf(c, "bucket.Set - Unable to get previous item, creating new\n%v", dErr)

		key = datastore.NewIncompleteKey(c, entity, nil)
		item.DataElement = repository.WithNew("site")
	} else {
		item.DataElement = repository.WithOld(oldItem.DataElement, "site")
	}

	key, err = datastore.Put(c, key, &item)
	return
}

//SetAll sets all items
//No logic is put in place for items with duplicate field,
//so field will be overwritten as it loops through.
//Latest value survives
//No transaction, so if any fail, it will not rollback any successful
func SetAll(c context.Context, newItems ...Item) (err error) {
	var (
		fields   []string
		needKeys []Item
		haveKeys []Item
	)

	for _, item := range newItems {
		fields = append(fields, item.Field)
	}

	oldItems, keys, _ := GetAll(c, fields...)

CheckingForNew:
	for _, newI := range newItems {
		for _, oldI := range oldItems {
			if newI.Field == oldI.Field {
				haveKeys = append(haveKeys, newI)
				continue CheckingForNew
			}
		}
		needKeys = append(needKeys, newI)
	}

	for _, needItem := range needKeys {
		keys = append(keys, datastore.NewIncompleteKey(c, entity, nil))
		needItem.DataElement = repository.WithNew("site")
		haveKeys = append(haveKeys, needItem)
	}

	_, err = datastore.PutMulti(c, keys, haveKeys)

	return
}

func Default(c context.Context, defaultItem Item) (appliedItem Item) {
	var items []Item

	q := datastore.NewQuery(entity).Filter("Field =", defaultItem.Field).Limit(1)

	_, err := q.GetAll(c, &items)

	if err != nil || len(items) > 0 {
		appliedItem = items[0]
		return
	}

	log.Infof(c, "Field %s missing, using default %s", defaultItem.Field, defaultItem.Value)
	key := datastore.NewIncompleteKey(c, entity, nil)
	defaultItem.DataElement = repository.WithNew("site")
	appliedItem = defaultItem

	datastore.Put(c, key, &appliedItem)

	return
}

func DefaultAll(c context.Context, defaultItems ...Item) (appliedItems []Item) {
	var (
		fields   []string
		needKeys []Item
		keys     []*datastore.Key
	)

	for _, item := range defaultItems {
		fields = append(fields, item.Field)
	}

	oldItems, _, _ := GetAll(c, fields...)

CheckingForNew:
	for _, defaultI := range defaultItems {
		for _, oldI := range oldItems {
			if defaultI.Field == oldI.Field {
				appliedItems = append(appliedItems, oldI)
				continue CheckingForNew
			}
		}
		needKeys = append(needKeys, defaultI)
	}

	for i := range needKeys {
		log.Infof(c, "Field %s missing, using default %s", needKeys[i].Field, needKeys[i].Value)
		keys = append(keys, datastore.NewIncompleteKey(c, entity, nil))
		needKeys[i].DataElement = repository.WithNew("site")
		appliedItems = append(appliedItems, needKeys[i])
	}

	_, err := datastore.PutMulti(c, keys, needKeys)
	if err != nil {
		log.Infof(c, "bucket.DefaultAll - Unable to put missing info into database\n%v", err)
	}

	return
}

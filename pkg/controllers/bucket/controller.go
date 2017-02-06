package bucket

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const entity = "bucket"

//Get gets an item from the bucket with the same field
func Get(c context.Context, field string) (item Item, err error) {
	var items []Item

	q := datastore.NewQuery(entity).Filter("Field =", field).Limit(1)

	_, err = q.GetAll(c, items)

	if err != nil {
		return
	}

	if len(items) == 0 {
		err = ErrFieldNotFound
	}

	return
}

//GetAll gets all items with the fields listed
func GetAll(c context.Context, fields ...string) (items []Item, err error) {
	for _, f := range fields {
		var item Item
		item, err = Get(c, f)
		if err != nil {
			return
		}
		items = append(items, item)
	}

	return
}

//Set sets the field with item
func Set(c context.Context, item Item) (err error) {
	var key *datastore.Key
	q := datastore.NewQuery(entity).Filter("Field =", item.Field).Limit(1)
	keys, dErr := q.GetAll(c, &item)
	if dErr != nil {
		log.Warningf(c, "bucket.Set - Unable to look into bucket\n%v", dErr)
	}

	if dErr != nil || len(keys) == 0 {
		key = datastore.NewIncompleteKey(c, entity, nil)
	} else {
		key = keys[0]
	}

	_, err = datastore.Put(c, key, &item)
	return
}

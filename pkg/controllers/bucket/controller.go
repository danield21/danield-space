package bucket

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

//Get gets an item from the bucket with the same field
func Get(context context.Context, field string) (item Item, err error) {
	var items []Item

	q := datastore.NewQuery("Bucket").Filter("Field =", field).Limit(1)

	_, err = q.GetAll(context, items)

	if err != nil {
		return
	}

	if len(items) == 0 {
		err = ErrFieldNotFound
	}

	return
}

//GetAll gets all items with the fields listed
func GetAll(context context.Context, fields ...string) (items []Item, err error) {
	for _, f := range fields {
		var item Item
		item, err = Get(context, f)
		if err != nil {
			return
		}
		items = append(items, item)
	}

	return
}

package bucket

import (
	"time"

	"github.com/danield21/danield-space/pkg/controllers"
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
	var items []Item
	var key *datastore.Key
	q := datastore.NewQuery(entity).Filter("Field =", item.Field).Limit(1)
	keys, dErr := q.GetAll(c, items)
	if dErr != nil {
		log.Warningf(c, "bucket.Set - Unable to look into bucket\n%v", dErr)
	}

	if dErr != nil || len(keys) == 0 {
		key = datastore.NewIncompleteKey(c, entity, nil)
		item.DataElement = controllers.DataElement{
			CreatedOn:  time.Now(),
			CreatedBy:  "site",
			ModifiedOn: time.Now(),
			ModifiedBy: "site",
		}
	} else {
		key = keys[0]
		oldItem := items[0]
		item.DataElement = controllers.DataElement{
			CreatedOn:  oldItem.CreatedOn,
			CreatedBy:  oldItem.CreatedBy,
			ModifiedOn: time.Now(),
			ModifiedBy: "site",
		}
	}

	_, err = datastore.Put(c, key, &item)
	return
}

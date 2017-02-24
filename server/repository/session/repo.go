package session

import (
	"time"

	"github.com/danield21/danield-space/server/repository"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

var entity = "Sessions"

func GetAll(ctx context.Context) ([]*Key, error) {
	var keys []*Key
	q := datastore.NewQuery(entity)
	dKeys, err := q.GetAll(ctx, &keys)

	if err != nil {
		return nil, err
	}

	for i := range keys {
		keys[i].Key = dKeys[i]
	}
	return keys, nil
}

func GetAllSince(ctx context.Context, t time.Time) ([]*Key, error) {
	var keys []*Key
	q := datastore.NewQuery(entity).Filter("CreatedOn >", t)
	dKeys, err := q.GetAll(ctx, &keys)

	if err != nil {
		return nil, err
	}

	for i := range keys {
		keys[i].Key = dKeys[i]
	}
	return keys, nil
}

func Put(ctx context.Context, key *Key) error {
	var err error

	key.DataElement = repository.WithNew("site")

	dKey := datastore.NewIncompleteKey(ctx, entity, nil)
	dKey, err = datastore.Put(ctx, dKey, key)
	if err == nil {
		key.Key = dKey
	}
	return err
}

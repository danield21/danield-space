package datastore

import (
	"time"

	"github.com/danield21/danield-space/server/store"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

var sessionEntity = "Sessions"

type Session struct{}

func (Session) GetAll(ctx context.Context) ([]*store.SessionKey, error) {
	var keys []*store.SessionKey
	q := datastore.NewQuery(sessionEntity)
	dKeys, err := q.GetAll(ctx, &keys)

	if err != nil {
		return nil, err
	}

	for i := range keys {
		keys[i].Key = dKeys[i]
	}
	return keys, nil
}

func (Session) GetAllSince(ctx context.Context, t time.Time) ([]*store.SessionKey, error) {
	var keys []*store.SessionKey
	q := datastore.NewQuery(sessionEntity).Filter("CreatedOn >", t)
	dKeys, err := q.GetAll(ctx, &keys)

	if err != nil {
		return nil, err
	}

	for i := range keys {
		keys[i].Key = dKeys[i]
	}
	return keys, nil
}

func (Session) Put(ctx context.Context, key *store.SessionKey) error {
	var err error

	key.DataElement = store.WithNew("site")

	dKey := datastore.NewIncompleteKey(ctx, sessionEntity, nil)
	dKey, err = datastore.Put(ctx, dKey, key)
	if err == nil {
		key.Key = dKey
	}
	return err
}

package account

import (
	"errors"

	"github.com/danield21/danield-space/server/repository"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const entity = "Admin"

var ErrNoMatch = errors.New("no user was found")
var ErrNilAccount = errors.New("account was nil")

func GetAll(ctx context.Context) ([]*Account, error) {
	var accounts []*Account
	q := datastore.NewQuery(entity)
	keys, err := q.GetAll(ctx, &accounts)

	if err != nil {
		return nil, err
	}

	for i := range keys {
		accounts[i].Key = keys[i]
	}

	return accounts, nil
}

func Get(ctx context.Context, username string) (*Account, error) {
	var accounts []*Account
	q := datastore.NewQuery(entity).Filter("Username = ", username)
	keys, err := q.GetAll(ctx, &accounts)

	if err != nil {
		return nil, err
	}

	if len(accounts) != 1 {
		return nil, ErrNoMatch
	}

	accounts[0].Key = keys[0]
	return accounts[0], nil
}

func Put(ctx context.Context, account *Account) error {
	if account == nil {
		return ErrNilAccount
	}

	oldAcct, err := Get(ctx, account.Username)
	if err != nil {
		account.DataElement = repository.WithNew("unknown")
		account.Key = datastore.NewIncompleteKey(ctx, entity, nil)
	} else {
		account.DataElement = repository.WithOld("unknown", oldAcct.DataElement)
	}

	account.Key, err = datastore.Put(ctx, account.Key, &account)
	return err
}

func CanLogIn(ctx context.Context, username string, password []byte) bool {
	accounts, err := GetAll(ctx)

	if err != nil || len(accounts) == 0 {
		log.Warningf(ctx, "admin.IsAdmin - Unable to retrieve Admin accounts from database, using default\n")
		accounts = append(accounts, &Default)
		Default.DataElement = repository.WithNew("site")
		Default.Key = datastore.NewIncompleteKey(ctx, entity, nil)
		_, err = datastore.Put(ctx, Default.Key, accounts[0])
		if err != nil {
			log.Warningf(ctx, "admin.IsAdmin - Unable to put default account into database\n%v", err)
		}
	}

	for _, acc := range accounts {
		if username == acc.Username && acc.SamePassword(password) {
			return true
		}
	}
	return false
}

func ChangePassword(ctx context.Context, username string, password []byte) error {
	var accounts []Account
	q := datastore.NewQuery(entity).Filter("Username =", username)
	key, err := q.GetAll(ctx, &accounts)

	if err != nil {
		return err
	}

	if len(accounts) != 1 {
		return errors.New("Unable to find account")
	}

	accounts[0].Password(password)

	_, err = datastore.Put(ctx, key[0], &accounts[0])
	return err
}

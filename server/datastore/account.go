package datastore

import (
	"errors"

	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const adminEntity = "Admin"

var ErrNoAccount = errors.New("no user was found")
var ErrNilAccount = errors.New("account was nil")

type Account struct{}

func (Account) GetAll(ctx context.Context) ([]*models.Account, error) {
	var accounts []*models.Account
	q := datastore.NewQuery(adminEntity)
	keys, err := q.GetAll(ctx, &accounts)

	if err != nil {
		return nil, err
	}

	for i := range keys {
		accounts[i].Key = keys[i]
	}

	return accounts, nil
}

func (Account) Get(ctx context.Context, username string) (*models.Account, error) {
	var accounts []*models.Account
	q := datastore.NewQuery(adminEntity).Filter("Username = ", username)
	keys, err := q.GetAll(ctx, &accounts)

	if err != nil {
		return nil, err
	}

	if len(accounts) != 1 {
		return nil, ErrNoAccount
	}

	accounts[0].Key = keys[0]
	return accounts[0], nil
}

func (a Account) Put(ctx context.Context, account *models.Account) error {
	if account == nil {
		return ErrNilAccount
	}

	oldAcct, err := a.Get(ctx, account.Username)
	if err != nil {
		account.DataElement = models.WithNew("unknown")
		account.Key = datastore.NewIncompleteKey(ctx, adminEntity, nil)
	} else {
		account.DataElement = models.WithOld("unknown", oldAcct.DataElement)
	}

	account.Key, err = datastore.Put(ctx, account.Key, account)
	return err
}

func (a Account) CanLogIn(ctx context.Context, username string, password []byte) bool {
	accounts, err := a.GetAll(ctx)

	if err != nil || len(accounts) == 0 {
		log.Warningf(ctx, "admin.IsAdmin - Unable to retrieve Admin accounts from database, using default\n")
		accounts = append(accounts, &models.DefaultAccount)
		models.DefaultAccount.DataElement = models.WithNew("site")
		models.DefaultAccount.Key = datastore.NewIncompleteKey(ctx, adminEntity, nil)
		_, err = datastore.Put(ctx, models.DefaultAccount.Key, accounts[0])
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

func (Account) ChangePassword(ctx context.Context, username string, password []byte) error {
	var accounts []models.Account
	q := datastore.NewQuery(adminEntity).Filter("Username =", username)
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

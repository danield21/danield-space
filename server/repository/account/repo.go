package account

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

//Default is the default account
//Be sure to change it
var Default = Account{Username: "Root", Hashword: nil}

const entity = "Admin"

func init() {
	Default.Password([]byte("ThisIsAVerySimplePassword!"))
}

func GetAll(ctx context.Context) ([]*Account, error) {
	var accounts []*Account
	q := datastore.NewQuery(entity)
	_, err := q.GetAll(ctx, &accounts)
	return accounts, err
}

func IsAdmin(ctx context.Context, username string, password []byte) bool {
	accounts, err := GetAll(ctx)

	if err != nil || len(accounts) == 0 {
		log.Warningf(ctx, "admin.IsAdmin - Unable to retrieve Admin accounts from database, using default\n")
		accounts = append(accounts, &Default)
		key := datastore.NewIncompleteKey(ctx, entity, nil)
		_, err = datastore.Put(ctx, key, accounts[0])
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

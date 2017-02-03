package admin

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

//Default is the default account
//Be sure to change it
var Default = Admin{Username: "Root", Hashword: nil}

func init() {
	Default.Password([]byte{
		'T', 'h', 'i', 's',
		'I', 's',
		'A',
		'V', 'e', 'r', 'y',
		'S', 'i', 'm', 'p', 'l', 'e',
		'P', 'a', 's', 's', 'w', 'o', 'r', 'd', '!',
	})
}

func GetAll(c context.Context) (admins []Admin, err error) {
	q := datastore.NewQuery("Admin")
	_, err = q.GetAll(c, admins)
	return
}

func IsAdmin(c context.Context, username string, password []byte) bool {
	var admins []Admin
	q := datastore.NewQuery("Admin").Filter("Username =", username)
	_, err := q.GetAll(c, admins)

	if err == nil {
		return len(admins) > 0 && admins[0].SamePassword(password)
	}

	log.Warningf(c, "admin.IsAdmin - Unable to retrieve Admin accounts from database, using default\n%v", err)

	key := datastore.NewIncompleteKey(c, "Admin", nil)
	_, err = datastore.Put(c, key, Default)

	if err != nil {
		log.Warningf(c, "admin.IsAdmin - Unable to put default account into database\n%v", err)
	}

	return Default.Username == username && Default.SamePassword(password)
}

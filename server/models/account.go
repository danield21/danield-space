package models

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type Account struct {
	DataElement
	Username string
	Hashword []byte
	Super    bool
}

//DefaultAccount is the default account
//Be sure to change it
var DefaultAccount = Account{Username: "Root", Hashword: nil, Super: true}

func init() {
	DefaultAccount.Password([]byte("ThisIsAVerySimplePassword!"))
}

var validUsername = regexp.MustCompile("^[a-zA-Z][\\w\\d]{5,}$")

func (a *Account) Password(password []byte) (err error) {
	a.Hashword, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return
}

func (a Account) SamePassword(password []byte) bool {
	err := bcrypt.CompareHashAndPassword(a.Hashword, password)
	return err == nil
}

func ClearPassword(password []byte) {
	for i := range password {
		password[i] = 0
	}
}

func ValidUsername(username string) bool {
	return validUsername.MatchString(username)
}

type AccountRepository interface {
	GetAll(ctx context.Context) ([]*Account, error)
	Get(ctx context.Context, username string) (*Account, error)
	Put(ctx context.Context, account *Account) error
	CanLogIn(ctx context.Context, username string, password []byte) bool
	ChangePassword(ctx context.Context, username string, password []byte) error
}

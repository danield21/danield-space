package account

import (
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Username string
	Hashword []byte
}

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

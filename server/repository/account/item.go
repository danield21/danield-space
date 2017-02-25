package account

import (
	"regexp"

	"github.com/danield21/danield-space/server/repository"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	repository.DataElement
	Username string
	Hashword []byte
	Super    bool
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

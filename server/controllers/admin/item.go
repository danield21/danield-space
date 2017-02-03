package admin

import "golang.org/x/crypto/bcrypt"

type Admin struct {
	Username string
	Hashword []byte
}

func (a Admin) Password(password []byte) (err error) {
	a.Hashword, err = bcrypt.GenerateFromPassword(password, bcrypt.MaxCost)
	return
}

func (a Admin) SamePassword(password []byte) bool {
	err := bcrypt.CompareHashAndPassword(a.Hashword, password)
	return err == nil
}

func ClearPassword(password []byte) {
	for i := range password {
		password[i] = 0
	}
}

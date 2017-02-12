package admin

import (
	"github.com/gorilla/sessions"
)

const UserKey = "user"

func GetUser(s *sessions.Session) (user string, signedIn bool) {
	iUser, ok := s.Values[UserKey]

	if !ok {
		return
	}

	user, ok = iUser.(string)
	if !ok {
		delete(s.Values, UserKey)
		return
	}

	signedIn = true

	return
}

func SetUser(s *sessions.Session, user string) {
	s.Values[UserKey] = user
}

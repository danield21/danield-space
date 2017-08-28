package session

import (
	"github.com/gorilla/sessions"
)

const UserKey = "user"

func User(s *sessions.Session) (string, bool) {
	if s == nil {
		return "", false
	}

	iUser, ok := s.Values[UserKey]

	if !ok {
		return "", false
	}

	user, ok := iUser.(string)
	if !ok {
		delete(s.Values, UserKey)
		return "", false
	}

	return user, true
}

func SetUser(s *sessions.Session, user string) {
	if s != nil {
		s.Values[UserKey] = user
	}
}

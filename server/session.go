package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

var sessionStore sessions.Store

var ErrCreateSession = errors.New("Unable to create session")

//GetSession returns a session using a secure key
func GetSession(ctx context.Context, e handler.Environment, r *http.Request) *sessions.Session {
	var err error

	if sessionStore == nil {
		sessionStore, err = NewStore(ctx, e)
		if err != nil {
			return nil
		}
	}

	session, err := sessionStore.Get(r, "danield-space")
	if err != nil {
		log.Warningf(ctx, "handler.GetSession - Unable to get session\n%v", err)
	}
	return session
}

func NewStore(ctx context.Context, e handler.Environment) (sessions.Store, error) {
	var kytes [][]byte
	keys, _ := e.Repository().Session().GetAllSince(ctx, time.Now().AddDate(0, 0, -3))

	if len(keys) == 0 {
		key, err := NewKeys()
		if err != nil {
			return nil, err
		}
		e.Repository().Session().Put(ctx, key)
		keys = append(keys, key)
	}

	for _, k := range keys {
		kytes = append(kytes, k.Hash, k.Block)
	}

	s := sessions.NewCookieStore(kytes...)

	s.Options.HttpOnly = true
	s.Options.MaxAge = 60 * 60 * 24

	return s, nil
}

func NewKeys() (*store.SessionKey, error) {
	hash := securecookie.GenerateRandomKey(64)
	block := securecookie.GenerateRandomKey(32)

	if hash == nil || block == nil {
		return nil, ErrCreateSession
	}

	key := new(store.SessionKey)
	*key = store.SessionKey{
		Hash:  hash,
		Block: block,
	}

	return key, nil
}

package server

import (
	"errors"
	"net/http"

	"time"

	"github.com/danield21/danield-space/server/repository/session"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var store sessions.Store

var ErrCreateSession = errors.New("Unable to session")

//GetSession returns a session using a secure key
func GetSession(r *http.Request) *sessions.Session {
	var err error
	ctx := appengine.NewContext(r)

	if store == nil {
		store, err = NewStore(ctx)
		if err != nil {
			return nil
		}
	}

	session, err := store.Get(r, "danield-space")
	if err != nil {
		log.Warningf(ctx, "handler.GetSession - Unable to get session\n%v", err)
	}
	return session
}

func NewStore(ctx context.Context) (sessions.Store, error) {
	var kytes [][]byte
	keys, _ := session.GetAllSince(ctx, time.Now().AddDate(0, 0, -3))

	if len(keys) == 0 {
		key, err := NewKeys()
		if err != nil {
			return nil, err
		}
		session.Put(ctx, key)
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

func NewKeys() (*session.Key, error) {
	hash := securecookie.GenerateRandomKey(64)
	block := securecookie.GenerateRandomKey(32)

	if hash == nil || block == nil {
		return nil, ErrCreateSession
	}

	key := new(session.Key)
	*key = session.Key{
		Hash:  hash,
		Block: block,
	}

	return key, nil
}

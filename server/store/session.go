package store

import (
	"context"
	"time"
)

type SessionKey struct {
	DataElement
	Hash  []byte
	Block []byte
}

type SessionRepository interface {
	GetAll(ctx context.Context) ([]*SessionKey, error)
	GetAllSince(ctx context.Context, t time.Time) ([]*SessionKey, error)
	Put(ctx context.Context, key *SessionKey) error
}

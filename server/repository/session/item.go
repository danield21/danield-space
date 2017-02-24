package session

import (
	"github.com/danield21/danield-space/server/repository"
)

type Key struct {
	repository.DataElement
	Hash  []byte
	Block []byte
}

package session

import (
	"github.com/danield21/danield-space/server/models"
)

type Key struct {
	models.DataElement
	Hash  []byte
	Block []byte
}

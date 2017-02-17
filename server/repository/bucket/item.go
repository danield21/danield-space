package bucket

import (
	"github.com/danield21/danield-space/server/repository"
)

type Item struct {
	repository.DataElement
	Field string
	Value string
	Type  string
}

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

func NewItem(Field, Value, Type string) *Item {
	item := new(Item)
	item.Field = Field
	item.Value = Value
	item.Type = Type
	return item
}

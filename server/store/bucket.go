package store

import (
	"context"
)

type BucketRepository interface {
	Get(ctx context.Context, field string) (*Item, error)
	GetAll(ctx context.Context, fields ...string) (have []*Item, missing []string)
	Set(ctx context.Context, item *Item) error
	SetAll(ctx context.Context, items ...*Item) error
	Default(ctx context.Context, defaultItem *Item) *Item
	DefaultAll(ctx context.Context, defaultItems ...*Item) []*Item
}

type Item struct {
	DataElement
	Field string
	Value string `datastore:",noindex"`
	Type  string
}

func NewItem(Field, Value, Type string) *Item {
	item := new(Item)
	item.Field = Field
	item.Value = Value
	item.Type = Type
	return item
}

package view

import (
	"errors"

	"golang.org/x/net/context"
)

var ErrNoPage = errors.New("context does not have page")

type uniqueKey string

const pageKey = uniqueKey("page")
const dataKey = uniqueKey("data")

func WithPage(ctx context.Context, page string) context.Context {
	return context.WithValue(ctx, pageKey, page)
}

func WithData(ctx context.Context, data interface{}) context.Context {
	return context.WithValue(ctx, dataKey, data)
}

func Page(ctx context.Context) (string, error) {
	iPage := ctx.Value(pageKey)
	page, ok := iPage.(string)
	if !ok {
		return "", ErrNoPage
	}
	return page, nil
}

func Data(ctx context.Context) interface{} {
	return ctx.Value(dataKey)
}

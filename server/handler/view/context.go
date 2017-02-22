package view

import (
	"errors"

	"golang.org/x/net/context"
)

var ErrNoPage = errors.New("context does not have page")
var ErrNoTheme = errors.New("context does not have theme")

type uniqueKey string

const pageKey = uniqueKey("page")
const themeKey = uniqueKey("theme")
const dataKey = uniqueKey("data")

func WithPage(ctx context.Context, page string) context.Context {
	return context.WithValue(ctx, pageKey, page)
}

func WithTheme(ctx context.Context, theme string) context.Context {
	return context.WithValue(ctx, themeKey, theme)
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

func Theme(ctx context.Context) (string, error) {
	iTheme := ctx.Value(themeKey)
	theme, ok := iTheme.(string)
	if !ok {
		return "", ErrNoTheme
	}
	return theme, nil
}

func Data(ctx context.Context) interface{} {
	return ctx.Value(dataKey)
}

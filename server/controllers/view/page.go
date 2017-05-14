package view

import (
	"encoding/json"
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
)

func HTMLHandler(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
	page, err := Page(ctx)
	if err != nil {
		return ctx, err
	}
	data := Data(ctx)

	return ctx, e.View(w, page, data)
}

func JSONHandler(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
	return ctx, json.NewEncoder(w).Encode(Data(ctx))
}

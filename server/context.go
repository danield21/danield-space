package server

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//ContextGenerater gets a new app engine context
type ContextGenerater struct{}

func (gen ContextGenerater) Generate(r *http.Request) (ctx context.Context) {
	return appengine.NewContext(r)
}

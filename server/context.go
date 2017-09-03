package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/handler"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//ContextGenerater gets a new app engine context
var ContextGenerater handler.ContextGeneratorFunc = func(r *http.Request) (ctx context.Context) {
	return appengine.NewContext(r)
}

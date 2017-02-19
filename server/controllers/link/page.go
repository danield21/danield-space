package link

import (
	"github.com/danield21/danield-space/server/service/view"
	"golang.org/x/net/context"
)

func PageContext(ctx context.Context, page string, data interface{}) context.Context {
	pCtx := view.WithPage(ctx, page)
	dCtx := view.WithData(pCtx, data)
	return dCtx
}

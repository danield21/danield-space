package status

import (
	"github.com/danield21/danield-space/server/handler"
)

func LinkAll(h handler.Handler) handler.Handler {
	return handler.Chain(h,
		CheckUnauthorizedLink,
		CheckNotFoundLink,
	)
}

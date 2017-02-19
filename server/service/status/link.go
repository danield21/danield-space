package status

import (
	"github.com/danield21/danield-space/server/service"
)

func LinkAll(h service.Handler) service.Handler {
	return service.Chain(h, UnauthorizedLink, NotFoundLink)
}

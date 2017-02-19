package admin

import (
	"github.com/danield21/danield-space/server/service"
)

type AdminModel struct {
	service.BaseModel
	User string
}

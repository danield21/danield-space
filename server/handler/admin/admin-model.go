package admin

import (
	"github.com/danield21/danield-space/server/handler"
)

type AdminModel struct {
	handler.BaseModel
	User string
}

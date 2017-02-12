package admin

import (
	"github.com/danield21/danield-space/pkg/handler"
)

type AdminModel struct {
	handler.BaseModel
	User string
}

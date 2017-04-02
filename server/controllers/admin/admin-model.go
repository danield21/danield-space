package admin

import "github.com/danield21/danield-space/server/controllers/view"

type AdminModel struct {
	view.BaseModel
	User string
}

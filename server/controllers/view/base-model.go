package view

import (
	"github.com/danield21/danield-space/server/repository/siteInfo"
)

//BaseModel is the minimun model for any page to work.
type BaseModel struct {
	SiteInfo siteInfo.SiteInfo `json:"-"`
}

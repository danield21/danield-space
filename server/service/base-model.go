package service

import (
	"github.com/danield21/danield-space/server/repository/siteInfo"
)

//BaseModel is the minimun model for any theme to work
//Not all theme have to use all of the fields,
//but themes cannot use any other fields besides these.
//This is to ensue that any theme will work
type BaseModel struct {
	SiteInfo siteInfo.SiteInfo
}

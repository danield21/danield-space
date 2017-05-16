package view

import "github.com/danield21/danield-space/server/models"

//BaseModel is the minimun model for any page to work.
type BaseModel struct {
	SiteInfo models.SiteInfo `json:"-"`
}

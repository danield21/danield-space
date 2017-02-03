package siteInfo

import (
	"errors"
)

//ErrSiteInfo appears when site is unable to get information about itself from the database
var ErrSiteInfo = errors.New("Unable to get site information")

package siteInfo

import (
	"github.com/danield21/danield-space/pkg/controllers/bucket"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//Default has default information about the site
var Default = SiteInfo{
	Title:       "Ballooneer's Code",
	Link:        "http://danield.space",
	Owner:       "Daniel J Dominguez",
	Description: "Sometimes, having a lofty head is necessary. This is a site is dedicated to having an overview discussion to code, without worrying too much about implementation. Of course, we will be touching the ground to get a better view of what may need to happen, but for the most part we will care about the overall look and feel.",
}

//Get gets all information about the site
func Get(c context.Context) (info SiteInfo, err error) {
	fields, dErr := bucket.GetAll(c, "Title", "Link", "Owner", "Description")

	if dErr != nil {
		err = ErrSiteInfo
		info = Default
		log.Warningf(c, "siteInfo.Get - Unable to get site information from database, switching to default")
		return
	}

	for _, f := range fields {
		switch f.Field {
		case "Title":
			info.Title = f.Value
		case "Link":
			info.Link = f.Value
		case "Owner":
			info.Owner = f.Value
		case "Description":
			info.Description = f.Value
		}
	}

	return
}

package siteInfo

import (
	"reflect"

	"github.com/danield21/danield-space/pkg/controllers/bucket"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

const bucketPrefix = "siteInfo-"
const titleField = bucketPrefix + "title"
const linkField = bucketPrefix + "link"
const ownerField = bucketPrefix + "owner"
const descriptionField = bucketPrefix + "description"

//Default has default information about the site
var Default = SiteInfo{
	Title:       "Ballooneer's Code",
	Link:        "http://danield.space",
	Owner:       "Daniel J Dominguez",
	Description: "Sometimes, having a lofty head is necessary. This is a site is dedicated to having an overview discussion to code, without worrying too much about implementation. Of course, we will be touching the ground to get a better view of what may need to happen, but for the most part we will care about the overall look and feel.",
}

//Get gets all information about the site
func Get(c context.Context) (info SiteInfo) {
	info = Default

	fields, err := bucket.GetAll(c,
		titleField,
		linkField,
		ownerField,
		descriptionField,
	)

	if err != nil {
		log.Warningf(c, "siteInfo.Get - Unable to get site information from database, switching to default")
		err = Set(c, Default)
		if err != nil {
			log.Warningf(c, "siteInfo.Get - Unable to set default site information")
		}
		return

	} else if len(fields) > reflect.TypeOf(info).NumField() {
		log.Warningf(c, "siteInfo.Get - Missing a few fields, will be using default values for those instead")
	}

	for _, f := range fields {
		switch f.Field {
		case titleField:
			info.Title = f.Value
		case linkField:
			info.Link = f.Value
		case ownerField:
			info.Owner = f.Value
		case descriptionField:
			info.Description = f.Value
		}
	}

	return
}

func Set(c context.Context, info SiteInfo) (err error) {
	items := []bucket.Item{
		bucket.Item{
			Field: titleField,
			Value: info.Title,
			Type:  "string",
		},
		bucket.Item{
			Field: linkField,
			Value: info.Link,
			Type:  "string",
		},
		bucket.Item{
			Field: ownerField,
			Value: info.Owner,
			Type:  "string",
		},
		bucket.Item{
			Field: descriptionField,
			Value: info.Description,
			Type:  "string",
		},
	}

	err = bucket.SetAll(c, items...)

	return
}

package datastore

import (
	"context"
	"errors"

	"github.com/danield21/danield-space/server/store"
)

const bucketPrefix = "siteInfo-"
const titleField = bucketPrefix + "title"
const linkField = bucketPrefix + "link"
const ownerField = bucketPrefix + "owner"
const descriptionField = bucketPrefix + "description"

//ErrSiteInfo appears when site is unable to get information about itself from the database
var ErrSiteInfo = errors.New("Unable to get site information")

type SiteInfo struct {
	Bucket store.BucketRepository
}

//Get gets all information about the site
func (s SiteInfo) Get(c context.Context) store.SiteInfo {
	items := siteInfoToItems(store.DefaultSiteInfo)

	fields := s.Bucket.DefaultAll(c, items...)

	return itemsToSiteInfo(fields)
}

func (s SiteInfo) Set(c context.Context, info store.SiteInfo) error {
	items := siteInfoToItems(info)

	err := s.Bucket.SetAll(c, items...)

	return err
}

func siteInfoToItems(info store.SiteInfo) []*store.Item {
	return []*store.Item{
		store.NewItem(titleField, info.Title, "string"),
		store.NewItem(linkField, info.Link, "string"),
		store.NewItem(ownerField, info.Owner, "string"),
		store.NewItem(descriptionField, info.Description, "string"),
	}
}

func itemsToSiteInfo(items []*store.Item) store.SiteInfo {
	var info store.SiteInfo
	for _, item := range items {
		switch item.Field {
		case titleField:
			info.Title = item.Value
		case linkField:
			info.Link = item.Value
		case ownerField:
			info.Owner = item.Value
		case descriptionField:
			info.Description = item.Value
		}
	}
	return info
}

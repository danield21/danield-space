package repository

import (
	"errors"

	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
)

const bucketPrefix = "siteInfo-"
const titleField = bucketPrefix + "title"
const linkField = bucketPrefix + "link"
const ownerField = bucketPrefix + "owner"
const descriptionField = bucketPrefix + "description"

//ErrSiteInfo appears when site is unable to get information about itself from the database
var ErrSiteInfo = errors.New("Unable to get site information")

type SiteInfo struct {
	Bucket models.BucketRepository
}

//Get gets all information about the site
func (s SiteInfo) Get(c context.Context) models.SiteInfo {
	items := siteInfoToItems(models.DefaultSiteInfo)

	fields := s.Bucket.DefaultAll(c, items...)

	return itemsToSiteInfo(fields)
}

func (s SiteInfo) Set(c context.Context, info models.SiteInfo) error {
	items := siteInfoToItems(info)

	err := s.Bucket.SetAll(c, items...)

	return err
}

func siteInfoToItems(info models.SiteInfo) []*models.Item {
	return []*models.Item{
		models.NewItem(titleField, info.Title, "string"),
		models.NewItem(linkField, info.Link, "string"),
		models.NewItem(ownerField, info.Owner, "string"),
		models.NewItem(descriptionField, info.Description, "string"),
	}
}

func itemsToSiteInfo(items []*models.Item) models.SiteInfo {
	var info models.SiteInfo
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

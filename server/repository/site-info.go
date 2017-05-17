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

//Default has default information about the site
var Default = models.SiteInfo{
	Title:       "Ballooneer's Code",
	Link:        "http://danield.space",
	Owner:       "Daniel J Dominguez",
	Description: "Sometimes, having a lofty head is necessary. This is a site is dedicated to having an overview discussion to code, without worrying too much about implementation. Of course, we will be touching the ground to get a better view of what may need to happen, but for the most part we will care about the overall look and feel.",
}

//Get gets all information about the site
func (s SiteInfo) Get(c context.Context) models.SiteInfo {
	items := siteInfoToItems(Default)

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

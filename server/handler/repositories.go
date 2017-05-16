package handler

import (
	"github.com/danield21/danield-space/server/models"
)

type Repositories interface {
	BucketRepository
	AboutRepository
	SiteInfoRepository
}

type BucketRepository interface {
	Bucket() models.BucketRepository
}

type AboutRepository interface {
	About() models.AboutRepository
}

type SiteInfoRepository interface {
	SiteInfo() models.SiteInfoRepository
}

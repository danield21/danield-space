package handler

import (
	"github.com/danield21/danield-space/server/models"
)

type Repositories interface {
	BucketRepository
	AboutRepository
	SiteInfoRepository
	SessionRepository
	AccountRepository
	ArticleRepository
	CategoryRepository
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

type SessionRepository interface {
	Session() models.SessionRepository
}

type AccountRepository interface {
	Account() models.AccountRepository
}

type ArticleRepository interface {
	Article() models.ArticleRepository
}

type CategoryRepository interface {
	Category() models.CategoryRepository
}

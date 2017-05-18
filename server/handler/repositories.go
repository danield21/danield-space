package handler

import (
	"github.com/danield21/danield-space/server/store"
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
	Bucket() store.BucketRepository
}

type AboutRepository interface {
	About() store.AboutRepository
}

type SiteInfoRepository interface {
	SiteInfo() store.SiteInfoRepository
}

type SessionRepository interface {
	Session() store.SessionRepository
}

type AccountRepository interface {
	Account() store.AccountRepository
}

type ArticleRepository interface {
	Article() store.ArticleRepository
}

type CategoryRepository interface {
	Category() store.CategoryRepository
}

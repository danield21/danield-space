package server

import (
	"github.com/danield21/danield-space/server/store/datastore"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
)

type RepositoryConnections struct {
	BucketDatastore   datastore.Bucket
	AboutDatastore    datastore.About
	SiteInfoDatastore datastore.SiteInfo
	SessionDatastore  datastore.Session
	AccountDatastore  datastore.Account
	ArticleDatastore  datastore.Article
	CategoryDatastore datastore.Category
}

func (rc RepositoryConnections) Bucket() store.BucketRepository {
	return rc.BucketDatastore
}

func (rc RepositoryConnections) About() store.AboutRepository {
	return rc.AboutDatastore
}

func (rc RepositoryConnections) SiteInfo() store.SiteInfoRepository {
	return rc.SiteInfoDatastore
}

func (rc RepositoryConnections) Session() store.SessionRepository {
	return rc.SessionDatastore
}

func (rc RepositoryConnections) Account() store.AccountRepository {
	return rc.AccountDatastore
}

func (rc RepositoryConnections) Article() store.ArticleRepository {
	return rc.ArticleDatastore
}

func (rc RepositoryConnections) Category() store.CategoryRepository {
	return rc.CategoryDatastore
}

func CreateRepository() handler.Repositories {
	connections := RepositoryConnections{}

	connections.BucketDatastore = datastore.Bucket{}
	connections.AboutDatastore = datastore.About{
		Bucket: connections.BucketDatastore,
	}
	connections.SiteInfoDatastore = datastore.SiteInfo{
		Bucket: connections.BucketDatastore,
	}
	connections.SessionDatastore = datastore.Session{}
	connections.AccountDatastore = datastore.Account{}
	connections.CategoryDatastore = datastore.Category{}
	connections.ArticleDatastore = datastore.Article{
		Category: connections.CategoryDatastore,
	}

	return connections
}

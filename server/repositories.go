package server

import (
	"github.com/danield21/danield-space/server/store/datastore"
)

type RepositoryConnections struct {
	Bucket   datastore.Bucket
	About    datastore.About
	SiteInfo datastore.SiteInfo
	Session  datastore.Session
	Article  datastore.Article
	Category datastore.Category
}

func CreateRepository() RepositoryConnections {
	connections := RepositoryConnections{}

	connections.Bucket = datastore.Bucket{}
	connections.About = datastore.About{
		Bucket: connections.Bucket,
	}
	connections.SiteInfo = datastore.SiteInfo{
		Bucket: connections.Bucket,
	}
	connections.Session = datastore.Session{}
	connections.Category = datastore.Category{}
	connections.Article = datastore.Article{
		Category: connections.Category,
	}

	return connections
}

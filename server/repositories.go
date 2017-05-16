package server

import (
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"github.com/danield21/danield-space/server/repository"
)

type RepositoryConnections struct {
	BucketDatastore repository.Bucket
	AboutDatastore  repository.About
}

func (rc RepositoryConnections) Bucket() models.BucketRepository {
	return rc.BucketDatastore
}

func (rc RepositoryConnections) About() models.AboutRepository {
	return rc.AboutDatastore
}

func CreateRepository() handler.Repositories {
	connections := RepositoryConnections{}

	connections.BucketDatastore = repository.Bucket{}
	connections.AboutDatastore = repository.About{
		Bucket: connections.BucketDatastore,
	}

	return connections
}

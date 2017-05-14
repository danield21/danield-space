package server

import (
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"github.com/danield21/danield-space/server/repository"
)

type RepositoryConnections struct {
	BucketDatastore repository.Bucket
}

func (rc RepositoryConnections) Bucket() models.BucketRepository {
	return rc.BucketDatastore
}

func CreateRepository() handler.RepositoryConnections {
	connections := RepositoryConnections{}

	connections.BucketDatastore = repository.Bucket{}

	return connections
}

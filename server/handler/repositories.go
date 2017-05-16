package handler

import (
	"github.com/danield21/danield-space/server/models"
)

type RepositoryConnections interface {
	BucketConnection
}

type BucketConnection interface {
	Bucket() models.BucketRepository
}

package handler

import (
	"github.com/danield21/danield-space/server/models"
)

type RepositoryConnections interface {
	Bucket() models.BucketRepository
}

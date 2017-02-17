package categories

import (
	"github.com/danield21/danield-space/pkg/repository"
)

type Category struct {
	repository.DataElement
	Title string
	Url   string
}

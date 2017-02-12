package bucket

import (
	"github.com/danield21/danield-space/pkg/controllers"
)

type Item struct {
	controllers.DataElement
	Field string
	Value string
	Type  string
}

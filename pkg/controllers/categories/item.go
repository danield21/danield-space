package categories

import (
	"github.com/danield21/danield-space/pkg/controllers"
)

type Category struct {
	controllers.DataElement
	Title string
	Url   string
}

package controllers

import (
	"time"
)

//DataElement contains basic informtion on data that has been retrieved by a database
type DataElement struct {
	CreatedOn  time.Time
	CreatedBy  string
	ModifiedOn time.Time
	ModifiedBy string
}
package bucket

import (
	"errors"
)

//ErrFieldNotFound appears when the field request does not exist in the database
var ErrFieldNotFound = errors.New("Field not found")

package bucket

import (
	"errors"
)

//ErrFieldNotFound appears when the field request does not exist in the database
var ErrFieldNotFound = errors.New("field not found")

//ErrNilItem appears when item parameter is nil
var ErrNilItem = errors.New("err was nil")

package form

import (
	"errors"
	"strings"
)

func NotEmpty(fld *Field, msg string) bool {
	if len(fld.Values) == 0 {
		Fail(fld, msg)
		return false
	} else if len(fld.Values) == 1 && strings.Trim(fld.Get(), " ") == "" {
		Fail(fld, msg)
		return false
	}
	return true
}

func Fail(fld *Field, msg string) {
	fld.Error = errors.New(msg)
}

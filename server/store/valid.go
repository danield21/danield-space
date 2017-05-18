package store

import (
	"regexp"
)

//ValidURLPart is a helper function to determine if a entered url can be valid
func ValidURLPart(urlPart string) bool {
	var valid = regexp.MustCompile("^([a-z]+(-[a-z]+)?)+$")
	return valid.MatchString(urlPart)
}

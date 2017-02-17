package controllers

import "regexp"

//ValidUrlPart is a helper function to determine if a entered url can be valid
func ValidUrlPart(urlPart string) bool {
	var valid = regexp.MustCompile("^([a-z]+(-[a-z]+)?)+$")
	return valid.MatchString(urlPart)
}

package form

import "strings"

func NotEmpty(f *Field, message string) bool {
	if strings.Trim(f.Value, " ") == "" {
		f.ErrorMessage = message
		return false
	}
	return true
}

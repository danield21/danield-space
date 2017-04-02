package form

import "strings"

func NotEmpty(f *Field, message string) bool {
	if strings.Trim(f.Value, " ") == "" {
		f.Error = true
		f.Message = message
		return false
	}
	return true
}

func Fail(f *Field, msg string) {
	f.Error = true
	f.Message = msg
}

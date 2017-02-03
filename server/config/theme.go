package config

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"regexp"
)

//DefaultTheme is the theme if no theme is specified
const DefaultTheme = "balloon"

//ErrInvalidTheme is the error returned when theme doesn't pass validation
var ErrInvalidTheme = errors.New("Theme only has letters and \"-\"")

//RenderTemplateWithTheme is a helper function to render golang templates with a theme
func RenderTemplateWithTheme(t *template.Template, w io.Writer, theme, view string, data interface{}) error {
	if !ValidTheme(theme) {
		return ErrInvalidTheme
	}

	var buffer = new(bytes.Buffer)

	err := t.ExecuteTemplate(buffer, "theme/"+theme+"/head", data)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buffer, view, data)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buffer, "theme/"+theme+"/footer", data)
	if err != nil {
		return err
	}

	_, err = buffer.WriteTo(w)

	return err

}

//ValidTheme is a helper function to determine if a entered theme can be valid
func ValidTheme(theme string) bool {
	var valid = regexp.MustCompile("^([a-z]+(-[a-z]+)?)+$")
	return valid.MatchString(theme)
}

//GetTheme gets the theme. If no theme was specified, then the default theme is given
func GetTheme(r *http.Request) (theme string) {
	err := r.ParseForm()
	if err != nil {
		return DefaultTheme
	}

	if theme, ok := r.Form["theme"]; ok {
		return theme[0]
	}

	return DefaultTheme
}

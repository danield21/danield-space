package envir

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/danield21/danield-space/pkg/repository/theme"
)

//RenderTemplateWithTheme is a helper function to render golang templates with a theme
func RenderTemplateWithTheme(t *template.Template, w io.Writer, useTheme, view string, data interface{}) error {
	if !theme.ValidTheme(useTheme) {
		return theme.ErrInvalidTheme
	}

	var buffer = new(bytes.Buffer)

	err := t.ExecuteTemplate(buffer, "theme/"+useTheme+"/head", data)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buffer, view, data)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buffer, "theme/"+useTheme+"/footer", data)
	if err != nil {
		return err
	}

	_, err = buffer.WriteTo(w)

	return err
}

//GetTheme gets the theme. If no theme was specified, then the default theme is given
func GetTheme(r *http.Request, defaultTheme string) string {
	err := r.ParseForm()
	if err != nil {
		return defaultTheme
	}

	if theme, ok := r.Form["theme"]; ok {
		return theme[0]
	}

	return defaultTheme
}

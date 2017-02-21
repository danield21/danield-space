package handler

import (
	"bytes"
	"html/template"
	"io"

	"github.com/danield21/danield-space/server/repository/theme"
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

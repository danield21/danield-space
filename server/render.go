package server

import (
	"bytes"
	"html/template"
	"io"
)

//RenderTemplate is a helper function to render golang templates
func RenderTemplate(t *template.Template, w io.Writer, view string, data interface{}) error {
	var buffer = new(bytes.Buffer)

	err := t.ExecuteTemplate(buffer, "core/head", data)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buffer, view, data)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(buffer, "core/footer", data)
	if err != nil {
		return err
	}

	_, err = buffer.WriteTo(w)

	return err
}

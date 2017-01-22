package config

import (
	"html/template"
	"io"
)

//MockConfig is only used for unit tests
type MockConfig struct {
	Templates *template.Template
}

//View creates a mock view
func (c MockConfig) View(w io.Writer, theme, view string, data interface{}) error {
	return RenderTemplateWithTheme(c.Templates, w, theme, view, data)
}

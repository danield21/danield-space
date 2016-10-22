package handlers

import (
	"html/template"
	"io"
)

//Config holds views and database connections to inject into handlers
type Config interface {
	View(w io.Writer, view string, data interface{}) error
}

//Only used for unit tests
type config struct {
	Templates *template.Template
}

func (c config) View(w io.Writer, view string, data interface{}) error {
	return c.Templates.ExecuteTemplate(w, view, data)
}

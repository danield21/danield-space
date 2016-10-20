package handlers

import "html/template"

//Config holds views and database connections to inject into handlers
type Config struct {
	Views *template.Template
}

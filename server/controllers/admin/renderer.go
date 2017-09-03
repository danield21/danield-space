package admin

import "io"

type Renderer interface {
	Render(w io.Writer, view string, data interface{}) error
	String(view string, data interface{}) (string, error)
}

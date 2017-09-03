package views

import (
	"bytes"
	"html/template"
	"io"
	"sync"
)

type Renderer struct {
	templates         *template.Template
	waitForView       sync.Mutex
	generateTemplates <-chan *template.Template
}

func (rnd *Renderer) waitToLoad() {
	rnd.waitForView.Lock()
	defer rnd.waitForView.Unlock()

	if rnd.templates == nil {
		rnd.templates = <-rnd.generateTemplates
	}
}

func (rnd *Renderer) Render(w io.Writer, view string, data interface{}) error {
	rnd.waitToLoad()
	return rnd.templates.ExecuteTemplate(w, view, data)
}

func (rnd *Renderer) String(view string, data interface{}) (string, error) {
	rnd.waitToLoad()

	buffer := new(bytes.Buffer)
	err := rnd.templates.ExecuteTemplate(buffer, view, data)

	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

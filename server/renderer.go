package server

import (
	"bytes"
	"html/template"
	"net/http"
	"sync"

	"github.com/danield21/danield-space/server/handler"
	"golang.org/x/net/context"
)

type Renderer struct {
	templates         *template.Template
	waitForView       sync.Mutex
	generateTemplates <-chan *template.Template
}

func NewRenderer(generator <-chan *template.Template) *Renderer {
	rnd := new(Renderer)

	rnd.generateTemplates = generator

	return rnd
}

func (rnd *Renderer) waitToLoad() {
	rnd.waitForView.Lock()
	defer rnd.waitForView.Unlock()

	if rnd.templates == nil {
		rnd.templates = <-rnd.generateTemplates
	}
}

func (rnd *Renderer) Render(ctx context.Context, view string, data interface{}) ([]byte, error) {
	rnd.waitToLoad()

	buffer := new(bytes.Buffer)
	err := rnd.templates.ExecuteTemplate(buffer, view, data)
	return buffer.Bytes(), err
}

func (rnd *Renderer) Send(w http.ResponseWriter, r *http.Request, pg *handler.Page) error {
	rnd.waitToLoad()

	if pg.Status != 0 {
		w.WriteHeader(pg.Status)
	}

	for header, value := range pg.Header {
		w.Header().Add(header, value)
	}

	return rnd.templates.ExecuteTemplate(w, "core/page", pg)
}

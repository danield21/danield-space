package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//Get gathers all the html files in the view directory and stores them in the Go template structure
func Get(directory string) <-chan *template.Template {
	out := make(chan *template.Template)
	views := template.New("view")

	funcs := template.FuncMap{
		"humanTime":   HumanTime,
		"machineTime": MachineTime,
		"asTitle":     strings.Title,
	}
	views.Funcs(funcs)

	go (func() {
		err := filepath.Walk(directory, addTo(views, funcs))
		if err != nil {
			log.Printf("ERROR: Error occured in initizing views: %v", err)
		}
		out <- views
	})()

	return out
}

func addTo(views *template.Template, funcs template.FuncMap) func(string, os.FileInfo, error) error {
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Printf("ERROR: Errror occured in walking to %s: %v", path, err)
			return err
		}

		if f.IsDir() {
			return nil
		}

		if err != nil {
			log.Printf("ERROR: Error occured in opening %s: %v", path, err)
			return err
		}

		contents, err := ioutil.ReadFile(path)

		if err != nil {
			log.Printf("ERROR: Error occured in reading %s: %v", path, err)
			return err
		}

		view, err := template.New("").Funcs(funcs).Parse(string(contents))

		if err != nil {
			log.Printf("ERROR: Error occured in parsing %s: %v", path, err)
			return err
		}

		viewName := getViewName(path)

		views.AddParseTree(viewName, view.Tree)
		log.Printf("INFO: Adding to view [%s]", viewName)

		return nil
	}
}

func getViewName(path string) string {
	normalizedPath := filepath.ToSlash(path)

	dir, file := filepath.Split(normalizedPath)

	dirs := strings.Split(dir, "/")
	withoutBase := strings.Join(dirs[1:], "/")

	parts := strings.Split(file, ".")
	name := strings.Join(parts[:len(parts)-1], ".")

	return withoutBase + name
}

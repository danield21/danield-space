package views

import (
	"html/template"
	"os"
	"path/filepath"
	"log"
	"strings"
)

func Get() *template.Template {
	views := template.New("view")
	err := filepath.Walk("view", addTo(views))
	if err != nil {
		log.Printf("[ERROR] Error occured in initizing views: %v", err)
	}
	return views
}

func addTo(views *template.Template) func(string, os.FileInfo, error) error {
	return func (path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Printf("[ERROR] Errror occured in walking to %s: %v", path, err)
			return err
		}

		if (f.IsDir()) {
			return nil
		}

		view, err := template.ParseFiles(path)

		if err != nil {
			log.Printf("[ERROR] Error occured in parsing %s: %v", path, err)
			return err
		}

		viewName := getViewName(path)

		views.AddParseTree(viewName, view.Tree)
		log.Printf("[INFO] Adding to view [%s]", viewName)

		return nil
	}
}

func getViewName(path string) string {
	normalizedPath := filepath.ToSlash(path)

	dir, file := filepath.Split(normalizedPath);

	dirs := strings.Split(dir, "/")
	withoutBase := strings.Join(dirs[1:], "/")

	parts := strings.Split(file, ".")
	name := strings.Join(parts[:len(parts) - 1], ".")

	return withoutBase + name
}
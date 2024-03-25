package main

import (
	"html/template"
	"io/fs"
	"jade-factory/go-snippetbox/internal/models"
	"jade-factory/go-snippetbox/ui"
	"path/filepath"
	"time"
)

// define a template data for hold structure for dynamic data.

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet

	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// function for formatting time
func humanDate(t time.Time) string {
	// Return the empty string if time has the zero value.
	if t.IsZero() {
		return ""
	}
	// Convert the time to UTC before formatting it.
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanData": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// init cache map
	cache := map[string]*template.Template{}

	// get a slice of all filepaths matching the template pattern
	// read from embedded filesystem
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	// iterate through the page
	for _, page := range pages {
		name := filepath.Base(page)

		// Create a slice containing the filepath patterns for the templates we // want to parse.
		patterns := []string{
			"html/base.tmpl.html", "html/partials/*.tmpl.html", page,
		}
		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

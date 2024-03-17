package main

import (
	"html/template"
	"jade-factory/go-snippetbox/internal/models"
	"path/filepath"
	"time"
)

// define a template data for hold structure for dynamic data.

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

// function for formatting time
func humanData(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04") // why this specific moment? oh It is just the numbers 1 2 3 4 5 6 7...
}

var functions = template.FuncMap{
	"humanData": humanData,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// init cache map
	cache := map[string]*template.Template{}

	// get a slice of all filepaths matching the template pattern
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	// iterate through the page
	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template file into a template set.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}
		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}
		// Call ParseFiles() *on this template set* to add the page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Add the template set to the map as normal...
		cache[name] = ts
	}
	return cache, nil
}

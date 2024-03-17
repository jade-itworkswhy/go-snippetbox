package main

import (
	"html/template"
	"jade-factory/go-snippetbox/internal/models"
	"path/filepath"
)

// define a template data for hold structure for dynamic data.

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
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

		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		// parse the files
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

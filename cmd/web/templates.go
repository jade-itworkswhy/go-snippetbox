package main

import "jade-factory/go-snippetbox/internal/models"

// define a template data for hold structure for dynamic data.

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

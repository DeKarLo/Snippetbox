package main

import "snippetbox.karen.kz/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

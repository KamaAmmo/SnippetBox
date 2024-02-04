package main

import (
	"html/template"
	"path/filepath"
	"snippetbox/internal/models"
)
type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}


func newTemplateCache() (map[string]*template.Template, error){

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil{
		return nil, err
	}

	for _, page := range pages{
		name := filepath.Base(page)

		files := []string{
			filepath.Join("./ui/html/base.tmpl.html"),
			filepath.Join("./ui/html/partials/nav.tmpl.html"),
			filepath.Join(page),
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil{
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
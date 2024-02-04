package main

import (
	"html/template"
	"path/filepath"
	"snippetbox/internal/models"
)
type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
	CurrentYear int
}


func newTemplateCache() (map[string]*template.Template, error){

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil{
		return nil, err
	}

	for _, page := range pages{
		name := filepath.Base(page)

	
		ts, err := template.ParseFiles(filepath.Join("./ui/html/base.tmpl.html"))
		if err != nil{
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join("./ui/html/partials/*.tmpl.html"))
		if err != nil{
			return nil, err
		}

		ts, err = ts.ParseFiles(filepath.Join(page))
		if err != nil{
			return nil, err
		}
	

		cache[name] = ts
	}

	return cache, nil
}
package main

import (
	// "fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"snippetbox/internal/models"
	"time"
)

type templateData struct {
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		Flash:       app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
	}
}

func humanDate(t time.Time) string {
	return t.Format(time.DateTime)
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts := template.New(name).Funcs(functions)

		ts, err := ts.ParseFiles(filepath.Join("./ui/html/base.tmpl.html"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join("./ui/html/partials/*.tmpl.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(filepath.Join(page))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

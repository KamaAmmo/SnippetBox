package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"snippetbox/internal/models"
	"time"
)


var functions = template.FuncMap{
	"humanDate": humanDate,
}

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
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
	}
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
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

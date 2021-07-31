package main

import (
	"SnippetBox/pkg/forms"
	"SnippetBox/pkg/models"
	"html/template"
	"path/filepath"
	"time"
)

type templateData struct {
	CurrentYear int
	Flash string
	Form *forms.Form
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error)  {
	//Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	//Extract the file name (like 'home,.page.html') from the full file path
	//the extension '.page.html'. this essentially give us the slice of all the path 'page' template
	//for the application
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	//Loop through the pages one-by-one
	for _, page := range pages {
		name := filepath.Base(page)
		//Parse the page template file in to a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//use the ParseGlob method to add any 'layout' template to
		//the template set (in our case it's just the 'base' layout at the moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		//use the ParseGlob method to add any 'partial' template to
		//the template set (in our case it's just the 'footer' layout at the moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string  {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate" : humanDate,
}
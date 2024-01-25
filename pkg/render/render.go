package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templateFile string) {
	templateCache, err := createTemplateCache()

	if err != nil {
		log.Fatal("Error creating template cache => ", err)
	}

	template, ok := templateCache[templateFile]

	if !ok {
		log.Fatal("Template not found => ", templateFile)
	}

	buffer := new(bytes.Buffer)

	err = template.Execute(buffer, nil)

	if err != nil {
		log.Fatal("Error executing template => ", err)
	}

	_, err = buffer.WriteTo(w)

	if err != nil {
		log.Fatal("Error writing template to browser => ", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		layoutMatches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(layoutMatches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = templateSet
	}

	return templateCache, nil
}

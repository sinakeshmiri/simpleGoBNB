package render

import (
	"bytes"
	"fmt"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/config"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var appConfig *config.Config

func NewTemplates(configurations *config.Config) {
	appConfig = configurations
}

var functions = template.FuncMap{}
var templateCache map[string]*template.Template

func RenderTemplate(writer http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	if appConfig.UseCache {
		templateCache = appConfig.TemplatesCache
	} else {
		var err error
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("!Failed")
		}

	}
	templ, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Failed")
	}
	myBuffer := new(bytes.Buffer)
	_ = templ.Execute(myBuffer, templateData)
	_, err := myBuffer.WriteTo(writer)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)
	var pages, err = filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}
		cache[name] = templateSet
	}
	return cache, nil
}

/*
	func RenderTemplate(writer http.ResponseWriter, tmpl string) {
		var parsedTemplate, _ = template.ParseFiles("./templates/" + tmpl)
		err := parsedTemplate.Execute(writer, nil)
		if err != nil {
			fmt.Println("error in rendering", err)
			return
		}
	}*/
/*func RenderTemplate(writer http.ResponseWriter, tmpl string) {

	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	templ, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}
	myBuffer := new(bytes.Buffer)
	_ = templ.Execute(myBuffer, nil)
	_, err = myBuffer.WriteTo(writer)
	if err != nil {
		fmt.Println(err)
	}

}*/

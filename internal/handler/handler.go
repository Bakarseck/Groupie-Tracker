package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

func render(w http.ResponseWriter, str string, value interface{}) {
	tmpl , err := TemplateCache(str)
	var 

}

func TemplateCache()(map[string]*template.Template, error) {
	cache := make(map[string]*template.Template
	pages,err := filepath.Glob("./templates/*.tmpl")
	if err != nil {
		log.Fatal("page not found")
	}
	for page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.Parsefiles(page))
		layout, err := filepath.Glob("/templates/layout/*.tmpl")
		if err != nil {
			log.Fatal("layout not found")
		}
		if len(layout)>0{
			tmpl.ParseGlob("/templates/layout/*.tmpl")
		}
		cache[name]=tmpl
	}
	return cache,nil
}

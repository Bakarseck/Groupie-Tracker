package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

func render(w http.ResponseWriter, tmpl string, value interface{}) {

}

func TemplateCache()(map[string]*template.Template, error) {
	cache := make(map[string]*template.Template
	pages,err := filepath.ParseGlob("./templates/*.tmpl")
	if err != nil {
		log.Fatal("page not found")
	}
	for page := range pages {
		
	}
}

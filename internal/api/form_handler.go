package api

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// GET /search
func SearchFormHandler(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	tmplPath := filepath.Join(dir, "templates", "search.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, nil)
}

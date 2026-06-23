package main

import (
	"net/http"
	"text/template"
)

const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
  <h1>{{.Title}}</h1>
  <p>{{.Body}}</p>
</body>
</html>
`

type PageData struct {
	Title string
	Body  string
}

func RenderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("page").Parse(tmplStr))
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")

	if title == "" || string(body) == "" {
		http.Error(w, "title and body are required", http.StatusBadRequest)
	}

	err := tmpl.Execute(w, PageData{Title: title, Body: string(body)})
	if err != nil {
		http.Error(w, "template execution failed", http.StatusInternalServerError)
	}
}

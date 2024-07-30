package test

import (
	"html/template"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("../main/templates/download.gohtml"))
		tmpl.ExecuteTemplate(writer, "download.gohtml", nil)
	})

	http.HandleFunc("/download", func(writer http.ResponseWriter, request *http.Request) {
		file := request.URL.Query().Get("file")
		if file != "" {
			writer.Header().Set("Content-Disposition", "attachment; filename=\""+file+"\"")
			http.ServeFile(writer, request, "./files/photo/"+file)
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}

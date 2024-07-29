package test

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestMultipleUploadViaAjax(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("../main/templates/upload_file_via_ajax.gohtml"))
			tmpl.ExecuteTemplate(writer, "view", nil)
		}
	})

	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		basePath, _ := os.Getwd()
		reader, err := request.MultipartReader()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			fileLocation := filepath.Join(basePath, "files", part.FileName())
			dst, err := os.Create(fileLocation)
			if dst != nil {
				defer dst.Close()
			}
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			if _, err := io.Copy(dst, part); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

		}
	})

	handlerStatisJQuery := http.FileServer(http.Dir("../main/assets/jquery/"))
	http.Handle("/static/jquery/", http.StripPrefix("/static/jquery/", handlerStatisJQuery))
	http.ListenAndServe("localhost:8080", nil)
}

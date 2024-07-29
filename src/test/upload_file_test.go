package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadFile(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("../main/templates/upload_file.gohtml"))
			err := tmpl.ExecuteTemplate(writer, "upload", nil)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(writer, "Only POST method is allowed", http.StatusMethodNotAllowed)
		}

		err := request.ParseMultipartForm(1024)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		alias := request.PostFormValue("alias")
		file, handler, err := request.FormFile("photo")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		filename := handler.Filename
		if alias != "" {
			filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
		}

		fileLocation := filepath.Join("files", "photo", filename)
		openFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0066)
		if err != nil {
			http.Error(writer, "cannot open file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer openFile.Close()

		_, err = io.Copy(openFile, file)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write([]byte("done"))
	})

	http.ListenAndServe("localhost:8080", nil)
}

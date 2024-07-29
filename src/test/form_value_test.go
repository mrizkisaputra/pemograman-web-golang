package test

import (
	"html/template"
	"net/http"
	"testing"
)

func TestFormValue(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("../main/templates/form_value.gohtml"))
			err := tmpl.ExecuteTemplate(writer, "form", nil)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}
		http.Error(writer, "", http.StatusBadRequest)
	})

	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("../main/templates/form_value.gohtml"))
		if request.Method == http.MethodPost {
			name := request.PostFormValue("name")
			password := request.PostFormValue("password")

			tmpl.ExecuteTemplate(
				writer,
				"result",
				map[string]any{
					"status":   "Register Success",
					"name":     name,
					"password": password,
				})
		}
		http.Error(writer, "", http.StatusBadRequest)
	})

	http.ListenAndServe("localhost:8080", nil)
}

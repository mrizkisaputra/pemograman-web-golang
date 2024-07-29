package test

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"testing"
)

func TestForm(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("../main/templates/form_ajax_json_payload.gohtml"))
			tmpl.ExecuteTemplate(writer, "form_json_payload", nil)
		}
	})

	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			decoder := json.NewDecoder(request.Body)
			payload := struct {
				Name   string `json:"name"`
				Age    int    `json:"age"`
				Gender string `json:"gender"`
			}{}
			if err := decoder.Decode(&payload); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			message := fmt.Sprintf(
				"hello, my name is %s. I'm %d year old %s",
				payload.Name,
				payload.Age,
				payload.Gender,
			)
			fmt.Println(decoder)
			writer.Write([]byte(message))
			return
		}
		http.Error(writer, "Only accept POST request", http.StatusBadRequest)

	})

	handlerStatisJS := http.FileServer(http.Dir("../main/assets/js/"))
	handlerStatisJQuery := http.FileServer(http.Dir("../main/assets/jquery/"))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", handlerStatisJS))
	http.Handle("/static/jquery/", http.StripPrefix("/static/jquery/", handlerStatisJQuery))

	http.ListenAndServe("localhost:8080", nil)
}

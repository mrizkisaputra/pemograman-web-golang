package test

import (
	"encoding/json"
	"net/http"
	"testing"
)

type User struct {
	Name  string
	Grade float32
}

func TestResponseJSON(t *testing.T) {
	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			users := []User{
				{Name: "Muhammat", Grade: 3.03},
				{Name: "Rizki", Grade: 3.50},
				{Name: "Muhammat Alim", Grade: 3.03},
				{Name: "Fajri Munawar", Grade: 3.20},
			}

			encode := json.NewEncoder(writer)
			err := encode.Encode(users)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
		}
		http.Error(writer, "", http.StatusBadRequest)
	})

	http.ListenAndServe("localhost:8080", nil)
}

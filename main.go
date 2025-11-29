package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {
	mux := http.NewServeMux()

	var todos []string

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {

		data, err := json.Marshal(todos)
		if err != nil {
			return
		}

		_, err = w.Write(data)

		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {
		var t TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, t.Item)
		w.WriteHeader(http.StatusCreated)

		_, err = w.Write([]byte("Todo created successfully!"))

		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

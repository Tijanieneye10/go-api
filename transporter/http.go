package transporter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tijanieneye10/go-api/internal/todo"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(svc *todo.Service) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(svc.GetAll())
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

		err = svc.Add(t.Item)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)

		_, err = w.Write([]byte("Todo created successfully!"))

		if err != nil {
			return
		}
	})

	return &Server{mux: mux}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}

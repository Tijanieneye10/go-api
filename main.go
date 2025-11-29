package main

import (
	"log"

	"github.com/Tijanieneye10/go-api/internal/todo"
	"github.com/Tijanieneye10/go-api/transporter"
)

func main() {
	svc := todo.NewService()

	server := transporter.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

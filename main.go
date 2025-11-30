package main

import (
	"log"

	"github.com/Tijanieneye10/go-api/internal/database"
	"github.com/Tijanieneye10/go-api/internal/todo"
	"github.com/Tijanieneye10/go-api/transporter"
)

func main() {
	db := database.NewDB()

	defer db.Sqlite.Close()

	svc := todo.NewService(db)

	server := transporter.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

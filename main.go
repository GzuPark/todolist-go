package main

import (
	"log"

	"github.com/gzupark/todolist-go/server"
)

func main() {
	if err := server.ListenAndServe(server.Config{
		Address:     ":8080",
		DatabaseURL: "postgres://postgres:password@localhost:5432/postgres?sslmode=disable",
	}); err != nil {
		log.Fatalln(err)
	}
}

package server

import (
	"net/http"

	"github.com/gzupark/todolist-go/api"
	"github.com/gzupark/todolist-go/db"
)

// Config server config
type Config struct {
	Address     string
	DatabaseURL string
}

// ListenAndServe database 접속 후 server 구동
func ListenAndServe(cfg Config) error {
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Address, loggingMiddleware(api.TodoListAPI()))
}

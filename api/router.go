package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TodoListAPI todo list API
func TodoListAPI() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)
	router.HandleFunc("/list", createTodoList).Methods(http.MethodPost)
	router.HandleFunc("/list/{list_id}", getTodoList).Methods(http.MethodGet)

	router.Use(handlePanic)

	return router
}

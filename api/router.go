package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TodoListAPI todo list API
func TodoListAPI() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)
	return router
}

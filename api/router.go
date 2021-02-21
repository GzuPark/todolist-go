package api

import (
	"fmt"
	"net/http"
)

// TodoListAPI welcome code
func TodoListAPI() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
}

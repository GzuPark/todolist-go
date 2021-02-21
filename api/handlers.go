package api

import (
	"net/http"

	"github.com/gzupark/todolist-go/db"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}

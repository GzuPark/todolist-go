package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gzupark/todolist-go/db"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	if err != nil {
		log.Println("internal error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{ Error string }{"Internal Error"})
		return
	}
	json.NewEncoder(w).Encode(lists)
}

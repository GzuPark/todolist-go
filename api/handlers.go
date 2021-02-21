package api

import (
	"net/http"

	"github.com/gzupark/todolist-go/db"
	"github.com/gzupark/todolist-go/todo"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}

func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	todoList, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, todoList)
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	list, err := db.GetTodoList(listID)
	must(err)
	writeJSON(w, list)
}

func updateTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	var req todo.List
	parseJSON(r.Body, &req)
	must(db.UpdateTodoList(listID, req.Name))
}

func deleteTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	must(db.DeleteTodoList(listID))
}

func createTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	var req todo.Item
	parseJSON(r.Body, &req)

	item, err := db.CreateTodoItem(listID, req.Text, req.Done)
	must(err)
	writeJSON(w, item)
}

func updateTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	var req todo.Item
	parseJSON(r.Body, &req)
	must(db.UpdateTodoItem(listID, itemID, req.Text, req.Done))
}

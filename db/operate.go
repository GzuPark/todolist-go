package db

import "github.com/gzupark/todolist-go/todo"

// GetTodoLists read all todo lists from the database
func GetTodoLists() ([]todo.List, error) {
	rows, err := db.Query(`SELECT id, name FROM todo_list`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := []todo.List{}
	for rows.Next() {
		var list todo.List
		if err := rows.Scan(&list.ID, &list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

// CreateTodoList create new todo list
func CreateTodoList(name string) (list todo.List, err error) {
	list.Name = name
	err = db.QueryRow(`INSERT INTO todo_list (name) VALUES ($1) RETURNING id`, name).Scan(&list.ID)
	return
}

// GetTodoList specific todo list with items
func GetTodoList(todoListID int) (todo.ListWithItems, error) {
	var list todo.ListWithItems

	rows, err := db.Query(`SELECT l.id, l.name, i.id, i.text, i.done
		FROM todo_list l
		LEFT JOIN todo_item i ON l.id = i.todo_list_id
		WHERE l.id = $1`, todoListID)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	list.Items = []todo.Item{}
	var gotTodoList bool
	for rows.Next() {
		var (
			itemID   *int
			itemText *string
			itemDone *bool
		)

		if err := rows.Scan(&list.ID, &list.Name, &itemID, &itemText, &itemDone); err != nil {
			return list, err
		}
		gotTodoList = true

		if itemID != nil && itemText != nil && itemDone != nil {
			list.Items = append(list.Items, todo.Item{
				ID:   *itemID,
				Text: *itemText,
				Done: *itemDone})
		}
	}

	if !gotTodoList {
		return list, ErrorNotFound
	}

	return list, nil
}

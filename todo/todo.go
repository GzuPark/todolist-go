package todo

// List todo list
type List struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Item todo list items
type Item struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// ListWithItems todo list with items
type ListWithItems struct {
	List
	Items []Item `json:"items"`
}

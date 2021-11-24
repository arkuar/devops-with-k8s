package db

type Todo struct {
	ID      int64
	Content string `json:"content"`
}

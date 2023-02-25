package models

const (
	Pending   string = "pending"
	Completed string = "completed"
)

type Task struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

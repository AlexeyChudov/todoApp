package Tasks

import "time"

type Task struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Completed    bool      `json:"completed"`
	CreationDate time.Time `json:"creation_date"`
	DueDate      time.Time `json:"due_date"`
	Priority     string    `json:"priority"`
}

type TaskFilter struct {
	Status   string
	Search   string
	FromDate time.Time
	ToDate   time.Time
	SortBy   string
	Order    string
	Limit    int
	Offset   int
}

type TaskService interface {
}

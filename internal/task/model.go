package task

import "time"

type Status string

const (
	StatusTodo       Status = "TODO"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

type Task struct {
	Id          string    `json:"id"`
	ProjectId   string    `json:"projectId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedBy   string    `json:"createdBy"`
	DueDate     time.Time `json:"dueDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskCreate struct {
	ProjectId   string    `json:"projectId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	CreatedBy   string    `json:"createdBy"`
}

type TaskUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

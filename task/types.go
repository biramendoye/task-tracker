package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` //todo, in-progress, done
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(id int, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type TaskList struct {
	Tasks []*Task `json:"tasks"`
}

func NewTaskList() *TaskList {
	return &TaskList{Tasks: make([]*Task, 0)}
}

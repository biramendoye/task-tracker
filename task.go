package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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

func (list *TaskList) insertTask(task *Task) {
	list.Tasks = append(list.Tasks, task)
}

func (list *TaskList) deleteTask(taskId int) bool {
	for i, task := range list.Tasks {
		if task.ID == taskId {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			return true
		}
	}

	return false
}

func (list *TaskList) updateTask(taskId int, description, status string) bool {
	for _, task := range list.Tasks {
		if task.ID == taskId {
			if description != "" {
				task.Description = description
			}
			if status != "" {
				task.Status = status
			}
			task.UpdatedAt = time.Now()
			return true
		}
	}

	return false
}

func (list *TaskList) GetAllTasks(status string) {
	fmt.Println("+----+----------------------+-------------+---------------------+---------------------+")
	fmt.Println("| ID | Description          | Status      | Created At          | Updated At          |")
	fmt.Println("+----+----------------------+-------------+---------------------+---------------------+")

	for _, task := range list.Tasks {
		if status == "" || task.Status == status {
			fmt.Printf("| %-2d | %-20s | %-11s | %-19s | %-19s |\n",
				task.ID,
				task.Description,
				task.Status,
				task.CreatedAt.Format("2006-01-02 15:04:05"),
				task.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
	fmt.Println("+----+----------------------+-------------+---------------------+---------------------+")
}

func (list *TaskList) saveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (list *TaskList) loadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s does not exist, starting with an empty task list.\n", filename)
			return nil
		}
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(list)
	if err != nil {
		return err
	}

	return nil
}

func (list *TaskList) size() int {
	return len(list.Tasks)
}

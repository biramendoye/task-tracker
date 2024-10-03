package task

import (
	"fmt"
	"time"
)

func (list *TaskList) InsertTask(task *Task) {
	list.Tasks = append(list.Tasks, task)
}

func (list *TaskList) DeleteTask(taskId int) bool {
	for i, task := range list.Tasks {
		if task.ID == taskId {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			return true
		}
	}

	return false
}

func (list *TaskList) UpdateTask(taskId int, description, status string) bool {
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

func (list *TaskList) Size() int {
	return len(list.Tasks)
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	storagePath := "tasks.json"

	action, description, id, status := parseArgs(args)

	// todos, err := LoadTodos(todoFile)
	taskList := NewTaskList()

	err := taskList.loadFromJSON(storagePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := taskList.saveToJSON(storagePath)
		if err != nil {
			fmt.Println("Error saving tasks to JSON:", err)
		}
	}()

	if action == "list" {
		taskList.GetAllTasks(status)
	} else if action == "add" {
		taskList.insertTask(NewTask(taskList.size()+1, description))
	} else if action == "delete" {
		taskList.deleteTask(id)
	} else if action == "update" {
		taskList.updateTask(id, description, status)
	} else if strings.HasPrefix(action, "mark-") {
		taskList.updateTask(id, description, strings.Replace(action, "mark-", "", -1))
	}
}

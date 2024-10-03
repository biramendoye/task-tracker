package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/biramendoye/task-tracker/task"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	storagePath := "tasks.json"

	action, description, id, status := parseArgs(args)

	taskList := task.NewTaskList()

	err := taskList.LoadFromJSON(storagePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := taskList.SaveToJSON(storagePath)
		if err != nil {
			fmt.Println("Error saving tasks to JSON:", err)
		}
	}()

	if action == "list" {
		taskList.GetAllTasks(status)
	} else if action == "add" {
		taskList.InsertTask(task.NewTask(taskList.Size()+1, description))
	} else if action == "delete" {
		taskList.DeleteTask(id)
	} else if action == "update" {
		taskList.UpdateTask(id, description, status)
	} else if strings.HasPrefix(action, "mark-") {
		taskList.UpdateTask(id, description, strings.Replace(action, "mark-", "", -1))
	}
}

func parseArgs(args []string) (string, string, int, string) {
	var (
		action      string
		description string
		status      string
		id          int
	)
	for i, arg := range args {
		if i == 0 {
			action = arg
		}
		if i == 1 && action == "add" {
			description = arg
		}

		if i == 1 && action == "list" {
			status = arg
		}
		if i == 1 && (action == "delete" || action == "update" || action == "mark-in-progress" || action == "mark-done") {
			id, _ = strconv.Atoi(arg)
		}
		if i == 2 && action == "update" {
			description = arg
		}
	}

	return action, description, id, status
}

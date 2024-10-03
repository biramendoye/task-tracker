package main

import "strconv"

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

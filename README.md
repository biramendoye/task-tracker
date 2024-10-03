# Task Tracker CLI

Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.

For detailed project instructions, please visit: [Task Tracker Project Instructions](https://roadmap.sh/projects/task-tracker).

## Features

- Add tasks with descriptions.
- Update tasks (description and status).
- Delete tasks.
- List tasks by status (e.g., todo, in-progress, done).
- Save and load tasks to/from a JSON file for persistence.

## Installation

### Prerequisites

- Go installed on your machine.
- The git command-line tool installed.

### Clone the Repository

```bash
git clone https://github.com/biramendoye/task-tracker.git
cd task-tracker
```

### Build the Project

```bash
go build -o task-cli cmd/main.go
```

## Usage

```bash
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Contributing

If you would like to contribute to this project, feel free to create a fork, make changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

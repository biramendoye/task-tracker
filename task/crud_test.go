package task

import "testing"

func TestInsertTask(t *testing.T) {
	taskList := NewTaskList()

	tests := []struct {
		name     string
		task     *Task
		expected int
	}{
		{"Insert Task 1", NewTask(1, "Task 1"), 1},
		{"Insert Task 2", NewTask(2, "Task 2"), 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			taskList.InsertTask(test.task)
			if len(taskList.Tasks) != test.expected {
				t.Errorf("Expected task list length %d, got %d", test.expected, len(taskList.Tasks))
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	taskList := NewTaskList()
	taskList.InsertTask(NewTask(1, "Task 1"))
	taskList.InsertTask(NewTask(2, "Task 2"))

	tests := []struct {
		name      string
		taskID    int
		expected  bool
		finalSize int
	}{
		{"Delete existing task", 1, true, 1},
		{"Delete non-existing task", 999, false, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := taskList.DeleteTask(test.taskID)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
			if len(taskList.Tasks) != test.finalSize {
				t.Errorf("Expected task list size %d, got %d", test.finalSize, len(taskList.Tasks))
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	taskList := NewTaskList()
	taskList.InsertTask(NewTask(1, "Task 1"))

	tests := []struct {
		name         string
		taskID       int
		description  string
		status       string
		expectedDesc string
		expectedStat string
		expectedRes  bool
	}{
		{"Update existing task", 1, "Updated Task", "", "Updated Task", "todo", true},
		{"Update existing task", 1, "", "in-progress", "Updated Task", "in-progress", true},
		{"Update non-existing task", 999, "New Desc", "todo", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := taskList.UpdateTask(test.taskID, test.description, test.status)
			if result != test.expectedRes {
				t.Errorf("Expected %v, got %v", test.expectedRes, result)
			}
			if test.expectedRes {
				updatedTask := taskList.Tasks[0]
				if updatedTask.Description != test.expectedDesc || updatedTask.Status != test.expectedStat {
					t.Errorf("Expected description %s and status %s, got %s and %s",
						test.expectedDesc, test.expectedStat, updatedTask.Description, updatedTask.Status)
				}
			}
		})
	}
}

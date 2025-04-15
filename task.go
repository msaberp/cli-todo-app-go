package main

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"due_date"`
	Priority  string    `json:"priority"`
}

// PriorityOrder returns a numeric value for priority comparison
func (t Task) PriorityOrder() int {
	switch t.Priority {
	case "High":
		return 3
	case "Medium":
		return 2
	case "Low":
		return 1
	default:
		return 0
	}
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil // no file yet
		}
		return nil, err
	}

	// If file is empty, return empty slice
	if len(file) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

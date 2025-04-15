package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
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

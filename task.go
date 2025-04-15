package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	file, err := ioutil.ReadFile(fileName)
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
	return ioutil.WriteFile(fileName, data, 0644)
}

// Function with no return values
func printTask(task Task) {
	fmt.Printf("Task: %s (ID: %d)\n", task.Title, task.ID)
}

// Function with one return value
func isTaskCompleted(task Task) bool {
	return task.Completed
}

// Function with multiple return values
func findTaskByID(tasks []Task, id int) (Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil // nil means no error
		}
	}
	return Task{}, errors.New("task not found")
}

// Function with named return values
func getTaskStats(tasks []Task) (total int, completed int) {
	for _, task := range tasks {
		total++
		if task.Completed {
			completed++
		}
	}
	return // returns total and completed
}
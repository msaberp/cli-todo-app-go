package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", -1, "Mark task as done (by ID)")
	del := flag.Int("del", -1, "Delete task (by ID)")
	flag.Parse()

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		t := Task{ID: len(tasks) + 1, Title: *add, Completed: false}
		tasks = append(tasks, t)
		fmt.Println("Added task:", t.Title)

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks.")
		}
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}

	case *done != -1:
		for i := range tasks {
			if tasks[i].ID == *done {
				tasks[i].Completed = true
				fmt.Println("Marked as done:", tasks[i].Title)
				break
			}
		}

	case *del != -1:
		for i := range tasks {
			if tasks[i].ID == *del {
				fmt.Println("Deleted:", tasks[i].Title)
				tasks = slices.Delete(tasks, i, i+1)
				break
			}
		}

		// reindexing the tasks
		for i, t := range tasks {
			t.ID = i + 1
			tasks[i] = t
		}

	default:
		fmt.Println("Invalid command. Use -add, -list, -done, or -del.")
	}

	saveTasks(tasks)
}

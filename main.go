package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	due := flag.String("due", "", "Due date (e.g., 2025-04-20)")
	priority := flag.String("priority", "Medium", "Task priority: Low, Medium, High")
	list := flag.Bool("list", false, "List all tasks")
	sort := flag.String("sort", "id", "Sort tasks by: id, due, priority")
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
		var dueDate time.Time
		var err error
		if *due != "" {
			dueDate, err = time.Parse("2006-01-02", *due)
			if err != nil {
				fmt.Println("Invalid date format. Use YYYY-MM-DD")
				os.Exit(1)
			}
		}
		t := Task{ID: len(tasks) + 1, Title: *add, Completed: false, DueDate: dueDate, Priority: *priority}
		tasks = append(tasks, t)
		fmt.Println("Added task:", t.Title)

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks.")
			break
		}

		switch *sort {
		case "id":
			slices.SortFunc(tasks, func(a, b Task) int { return a.ID - b.ID })
		case "due":
			slices.SortFunc(tasks, func(a, b Task) int {
				if a.DueDate.IsZero() && b.DueDate.IsZero() {
					return 0
				}
				if a.DueDate.IsZero() {
					return 1
				}
				if b.DueDate.IsZero() {
					return -1
				}
				return int(a.DueDate.Sub(b.DueDate))
			})
		case "priority":
			slices.SortFunc(tasks, func(a, b Task) int {
				return b.PriorityOrder() - a.PriorityOrder()
			})
		}

		// Print table header
		fmt.Printf("%-3s %-1s %-6s %-12s %-50s %-10s\n", "ID", "✓", "Priority", "Due Date", "Title", "Status")
		fmt.Println("----------------------------------------------------------------------------------------")

		// Print tasks
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "✓"
			}
			dueDateStr := "No date"
			if !t.DueDate.IsZero() {
				dueDateStr = t.DueDate.Format("2006-01-02")
			}
			fmt.Printf("%-3d %-1s %-6s %-12s %-50s %-10s\n",
				t.ID,
				status,
				t.Priority,
				dueDateStr,
				t.Title,
				func() string {
					if t.Completed {
						return "Completed"
					}
					return "Pending"
				}(),
			)
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

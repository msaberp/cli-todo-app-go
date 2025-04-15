# CLI Todo App in Go

A simple command-line todo application written in Go that helps you manage your tasks efficiently.

## Features

- Add new tasks with due dates and priorities
- List all tasks in a nicely formatted table
- Sort tasks by ID, due date, or priority
- Mark tasks as completed
- Delete tasks
- Persistent storage using JSON file
- Clean and intuitive CLI interface

## Installation

1. Make sure you have Go installed on your system
2. Clone this repository:
   ```bash
   git clone https://github.com/msaberp/cli-todo-app-go.git
   cd cli-todo-app-go
   ```
3. Initialize the Go module:
   ```bash
   go mod init github.com/msaberp/cli-todo-app-go
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build
   ```

## Usage

The application provides the following commands:

- Add a new task:
  ```bash
  ./cli-todo-app-go -add "Your task description" -due "2024-04-20" -priority "High"
  ```

- List all tasks:
  ```bash
  ./cli-todo-app-go -list
  ```

- List tasks sorted by due date:
  ```bash
  ./cli-todo-app-go -list -sort due
  ```

- List tasks sorted by priority:
  ```bash
  ./cli-todo-app-go -list -sort priority
  ```

- Mark a task as completed (using task ID):
  ```bash
  ./cli-todo-app-go -done <task-id>
  ```

- Delete a task (using task ID):
  ```bash
  ./cli-todo-app-go -del <task-id>
  ```

## Task Display Format

Tasks are displayed in a tabular format with the following columns:
- ID: Task identifier
- ✓: Completion status (✓ for completed, blank for pending)
- Priority: Task priority (High, Medium, Low)
- Due Date: Task due date in YYYY-MM-DD format
- Title: Task description
- Status: Task status (Completed or Pending)

Example output:
```
ID  ✓ Priority Due Date     Title                                              Status    
----------------------------------------------------------------------------------------
1   ✓ High     2024-04-20   Complete project documentation                     Completed
2     Medium   2024-04-25   Review pull requests                               Pending   
3     Low      No date      Update README                                      Pending   
```

## Data Storage

Tasks are stored in a `tasks.json` file in the application directory. The file is automatically created when you add your first task.

## Project Structure

- `main.go`: Contains the main application logic and CLI interface
- `task.go`: Defines the Task struct and related functions for task management

## License

This project is open source and available under the MIT License.


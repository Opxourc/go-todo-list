package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"todo-list/models"
)

const dateFormat = "2006-01-02"

// Run starts the interactive todo-list session. Tasks live only for this run.
func Run() {
	tasks := make([]models.Task, 0)
	nextID := 1

	printPrompt := func() {
		fmt.Print("> ")
	}

	printPrompt()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			printPrompt()
			continue
		}

		arguments := strings.Fields(input)
		command := strings.ToLower(arguments[0])
		arguments = arguments[1:]

		switch command {
		case "add":
			task, err := parseNewTask(arguments, nextID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				tasks = append(tasks, task)
				nextID++
				fmt.Printf("Added task %d.\n", task.ID)
			}
		case "list":
			listTasks(tasks)
		case "complete":
			id, err := parseID(arguments, "complete <id>")
			if err != nil {
				fmt.Println("Error:", err)
			} else if taskIndex := findTask(tasks, id); taskIndex == -1 {
				fmt.Printf("Task %d not found.\n", id)
			} else {
				tasks[taskIndex].Complete = true
				fmt.Printf("Completed task %d.\n", id)
			}
		case "delete":
			id, err := parseID(arguments, "delete <id>")
			if err != nil {
				fmt.Println("Error:", err)
			} else if taskIndex := findTask(tasks, id); taskIndex == -1 {
				fmt.Printf("Task %d not found.\n", id)
			} else {
				tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
				fmt.Printf("Deleted task %d.\n", id)
			}
		case "edit":
			if len(arguments) == 0 {
				fmt.Println("Error: usage: edit <id> <due-date> <objective>")
				break
			}
			id, err := parseID(arguments[:1], "edit <id> <due-date> <objective>")
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			taskIndex := findTask(tasks, id)
			if taskIndex == -1 {
				fmt.Printf("Task %d not found.\n", id)
				break
			}
			updatedTask, err := parseNewTask(arguments[1:], id)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			updatedTask.Complete = tasks[taskIndex].Complete
			tasks[taskIndex] = updatedTask
			fmt.Printf("Updated task %d.\n", id)
		case "help":
			RunHelp()
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Printf("Unknown command %q. Type 'help' for available commands.\n", command)
		}

		printPrompt()
	}
}

func parseNewTask(arguments []string, id int) (models.Task, error) {
	if len(arguments) < 2 {
		return models.Task{}, fmt.Errorf("usage: add <due-date> <objective>")
	}

	dueDate, err := time.Parse(dateFormat, arguments[0])
	if err != nil {
		return models.Task{}, fmt.Errorf("due date must use YYYY-MM-DD format")
	}

	return models.Task{
		ID:        id,
		DueDate:   dueDate,
		Objective: strings.Join(arguments[1:], " "),
	}, nil
}

func parseID(arguments []string, usage string) (int, error) {
	if len(arguments) != 1 {
		return 0, fmt.Errorf("usage: %s", usage)
	}

	id, err := strconv.Atoi(arguments[0])
	if err != nil || id < 1 {
		return 0, fmt.Errorf("task ID must be a positive number")
	}
	return id, nil
}

func findTask(tasks []models.Task, id int) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}
	return -1
}

func listTasks(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for _, task := range tasks {
		status := "pending"
		if task.Complete {
			status = "complete"
		}
		fmt.Printf("%d | %s | %s | %s\n", task.ID, task.DueDate.Format(dateFormat), status, task.Objective)
	}
}

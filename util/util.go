package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"todo-list/models"
)

const DateFormat = "2006-01-02" // Other files need to see this

// Parse a requested new task.
// The return is the task object and a error, if there's any.
func ParseNewTask(arguments []string, id int) (models.Task, error) {
	if len(arguments) < 2 {
		return models.Task{}, fmt.Errorf("usage: add <due-date> <objective>")
	}

	dueDate, err := time.Parse(DateFormat, arguments[0])
	if err != nil {
		return models.Task{}, fmt.Errorf("due date must use YYYY-MM-DD format")
	}

	return models.Task{
		ID:        id,
		DueDate:   dueDate,
		Objective: strings.Join(arguments[1:], " "),
	}, nil
}

// Parse an checking if it's it can be properly read and eventually used.
func ParseID(arguments []string, usage string) (int, error) {
	if len(arguments) != 1 {
		return 0, fmt.Errorf("usage: %s", usage)
	}

	id, err := strconv.Atoi(arguments[0])
	if err != nil || id < 1 {
		return 0, fmt.Errorf("task ID must be a positive number")
	}
	return id, nil
}

// Find a task with the matching ID in the provide slice of tasks.
// If the index of the task is not found, -1 will be returned instead.
func FindTask(tasks []models.Task, id int) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}
	return -1
}

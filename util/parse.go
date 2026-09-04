package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"todo-list/models"
)

const dateFormat = "2006-01-02"

// Parse a requested new task.
// The return is the task object and a error, if there's any.
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

// Parse an checking if it's it can be properly read and eventually used.
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

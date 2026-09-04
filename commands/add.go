package commands

import (
	"fmt"
	"todo-list/database"
	"todo-list/util"
)

// Add a new task to the memory table.
// The next ID is return since it gets incremeneted.
func Add(arguments []string, nextID int) int {
	task, err := util.ParseNewTask(arguments, nextID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		database.Tasks = append(database.Tasks, task)
		nextID++
		fmt.Printf("Added task %d.\n", task.ID)
	}
	return nextID
}

package commands

import (
	"fmt"
	"todo-list/database"
	"todo-list/util"
)

// Set the completion status of a task.
func Complete(arguments []string) {
	id, err := util.ParseID(arguments, "complete <id>")
	if err != nil {
		fmt.Println("Error: ", err)
	} else if taskIndex := util.FindTask(database.Tasks, id); taskIndex == -1 {
		fmt.Printf("Task %d not found.\n", id)
	} else {
		database.Tasks[taskIndex].Complete = true
		fmt.Printf("Completed task %d.\n", id)
	}
}

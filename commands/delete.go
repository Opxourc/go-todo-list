package commands

import (
	"fmt"
	"todo-list/database"
	"todo-list/util"
)

// Delete and remove a task from the memory table.
func Delete(arguments []string) {
	id, err := util.ParseID(arguments, "delete <id>")
	if err != nil {
		fmt.Println("Error:", err)
	} else if taskIndex := util.FindTask(database.Tasks, id); taskIndex == -1 {
		fmt.Printf("Task %d not found.\n", id)
	} else {
		database.Tasks = append(
			database.Tasks[:taskIndex],
			database.Tasks[taskIndex+1:]...,
		)
		fmt.Printf("Deleted task %d.\n", id)
	}
}

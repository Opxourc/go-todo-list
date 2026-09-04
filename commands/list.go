package commands

import (
	"fmt"
	"todo-list/database"
	"todo-list/util"
)

// Print all tasks that are in memory in a list.
func List() {
	if len(database.Tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for _, task := range database.Tasks {
		status := "pending"
		if task.Complete {
			status = "complete"
		}
		fmt.Printf("%d | %s | %s | %s\n",
			task.ID,
			task.DueDate.Format(util.DateFormat),
			status,
			task.Objective,
		)
	}
}

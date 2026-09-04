package commands

import (
	"fmt"
	"todo-list/database"
	"todo-list/util"
)

// Edit an already existing task.
func Edit(arguments []string) {
	if len(arguments) == 0 {
		fmt.Println("Error: usage: edit <id> <due-date> <objective>")
		return
	}
	id, err := util.ParseID(arguments[:1], "edit <id> <due-date> <objective>")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	taskIndex := util.FindTask(database.Tasks, id)
	if taskIndex == -1 {
		fmt.Printf("Task %d not found.\n", id)
		return
	}
	updatedTask, err := util.ParseNewTask(arguments[1:], id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	updatedTask.Complete = database.Tasks[taskIndex].Complete
	database.Tasks[taskIndex] = updatedTask
	fmt.Printf("Updated task %d.\n", id)
}

package commands

import "fmt"

// Print all commands that can be used by the user.
func Help() {
	fmt.Println("Usage: <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add <due-date> <objective>          Add a task")
	fmt.Println("  list                                List all tasks")
	fmt.Println("  complete <id>                       Mark a task as complete")
	fmt.Println("  delete <id>                         Delete a task")
	fmt.Println("  edit <id> <due-date> <objective>    Edit a task")
	fmt.Println("  help                                Show this help message")
	fmt.Println("  quit                                Exit the application")
	fmt.Println()
	fmt.Println("Dates must use YYYY-MM-DD format.")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-list/commands"
	"todo-list/database"
	"todo-list/models"
)

// Prints the starting menu when the program starts.
func printStart() {
	fmt.Println("--- Todo List ---")
	fmt.Println("Manage tasks that need to be completed!")
	fmt.Println("Type 'help' to get a list for a list of commands.")
}

func printPrompt() {
	fmt.Print("> ")
}

func main() {
	// Print and get all saved tasks if any
	printStart()
	database.GetSavedTasks()
	defer database.SaveTasksToFile()

	// Set up the ID's for the tasks
	nextID := 1
	for _, task := range database.Tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}

	// Use a scanner to clear bad inputs
	printPrompt()
	scanner := bufio.NewScanner(os.Stdin)

	// Keep grabbing inputs until the user quits
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" { // If empty, then just ignore and continue
			printPrompt()
			continue
		}

		// Commands will have arguements to them, grab them and what command was used
		fields := strings.Fields(input)
		command := models.Command(strings.ToLower(fields[0])) // Command is always the 0 index
		arguments := fields[1:]                               // 1 and up indexes

		// Define action based on the command
		switch command {
		case models.CommandAdd:
			nextID = commands.Add(arguments, nextID)
		case models.CommandList:
			commands.List()
		case models.CommandComplete:
			commands.Complete(arguments)
		case models.CommandDelete:
			commands.Delete(arguments)
		case models.CommandEdit:
			commands.Edit(arguments)
		case models.CommandHelp:
			commands.Help()
		case models.CommandQuit:
			return
		default:
			fmt.Printf("Unknown command %q. Type 'help' for avaliable commands.\n",
				command,
			)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

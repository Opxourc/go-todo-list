package main

import (
	"fmt"

	"todo-list/internal"
)

// Prints the starting menu when the program starts.
func printStart() {
	fmt.Println("--- Todo List ---")
	fmt.Println("Manage tasks that need to be completed!")
	fmt.Println("Type 'help' to get a list for a list of commands.")
}

func main() {
	printStart()
	internal.RunHelp()
	internal.Run()
}

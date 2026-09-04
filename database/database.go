package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todo-list/models"
)

// Where all tasks will be stored in memory.
var Tasks = make([]models.Task, 0)

const FILE_PATH = "data/tasks.json" // Where all tasks will be stored for presistence

// Attempts to open the file.
// If one doesn't exist, a new one is created.
// Returns the pointer to the file if successful.
func openFile() *os.File {
	file, openFileErr := os.OpenFile(
		FILE_PATH,
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if openFileErr != nil {
		log.Fatalf("Failed to open or create %v: %v\n",
			FILE_PATH,
			openFileErr,
		)
	}
	return file
}

// Retrieves the file that stored tasks for presistence and stores it in memory.
func GetSavedTasks() {
	// Check if the file exists, if not create a new one
	file := openFile()
	defer file.Close()

	// Read the contents of the file
	fileContents, readFileErr := os.ReadFile(FILE_PATH)
	if readFileErr != nil {
		log.Fatalf("Failed to read %v: %v\n",
			FILE_PATH,
			readFileErr,
		)
	}

	// Check if there's any bytes (any json string value)
	if len(bytes.TrimSpace(fileContents)) == 0 {
		return // Slice is already empty from initialization
	}

	// Convert from byte to Task objects if possible
	unmarshalErr := json.Unmarshal(fileContents, &Tasks)
	if unmarshalErr != nil {
		log.Fatalf("Failed to parse %v: %v\n", FILE_PATH, unmarshalErr)
	}
}

// Saves all tasks in memory to a file on disk for presistent storage.
func SaveTasksToFile() {
	// Check if the file exists, if not create a new one
	file, openFileErr := os.OpenFile(
		FILE_PATH,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if openFileErr != nil {
		log.Fatalf("Failed to open or create %v: %v\n",
			FILE_PATH,
			openFileErr,
		)
	}
	defer file.Close()

	// Parse the tasks in memory to a JSON string
	jsonString, jsonMarshalErr := json.MarshalIndent(Tasks, "", "	")
	if jsonMarshalErr != nil {
		log.Fatalf("Failed to parse to JSON: %v\n",
			jsonMarshalErr,
		)
	}

	// Write to the file with the JSON string
	_, fileWriteErr := file.Write(jsonString)
	if fileWriteErr != nil {
		log.Fatalf("Failed to save to %v: %v\n",
			FILE_PATH,
			fileWriteErr,
		)
	}

	fmt.Println("Successfully saved all tasks!")
}

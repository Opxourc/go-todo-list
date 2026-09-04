package database

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"todo-list/models"
)

func useTemporaryDataFile(t *testing.T) {
	t.Helper()

	dataDirectory := filepath.Dir(FILE_PATH)
	if err := os.MkdirAll(dataDirectory, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(FILE_PATH, []byte(" \n\t"), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGetSavedTasksIgnoresWhitespaceOnlyFile(t *testing.T) {
	workingDirectory := t.TempDir()
	originalDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalDirectory) })
	if err := os.Chdir(workingDirectory); err != nil {
		t.Fatal(err)
	}
	useTemporaryDataFile(t)

	Tasks = []models.Task{}
	GetSavedTasks()

	if len(Tasks) != 0 {
		t.Fatalf("expected no tasks, got %d", len(Tasks))
	}
}

func TestGetSavedTasksLoadsJSON(t *testing.T) {
	workingDirectory := t.TempDir()
	originalDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalDirectory) })
	if err := os.Chdir(workingDirectory); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(FILE_PATH), 0755); err != nil {
		t.Fatal(err)
	}
	contents := []byte(`[{"ID":7,"DueDate":"0001-01-01T00:00:00Z","Objective":"loaded","Complete":false}]`)
	if err := os.WriteFile(FILE_PATH, contents, 0644); err != nil {
		t.Fatal(err)
	}

	Tasks = []models.Task{}
	GetSavedTasks()

	if len(Tasks) != 1 || Tasks[0].ID != 7 || Tasks[0].Objective != "loaded" {
		t.Fatalf("unexpected loaded tasks: %#v", Tasks)
	}
}

func TestSaveTasksToFileTruncatesExistingJSON(t *testing.T) {
	workingDirectory := t.TempDir()
	originalDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalDirectory) })
	if err := os.Chdir(workingDirectory); err != nil {
		t.Fatal(err)
	}
	useTemporaryDataFile(t)

	Tasks = []models.Task{{ID: 1, Objective: "short"}}
	SaveTasksToFile()

	contents, err := os.ReadFile(FILE_PATH)
	if err != nil {
		t.Fatal(err)
	}
	var savedTasks []models.Task
	if err := json.Unmarshal(contents, &savedTasks); err != nil {
		t.Fatalf("saved JSON is invalid: %v", err)
	}
	if len(savedTasks) != 1 || savedTasks[0].Objective != "short" {
		t.Fatalf("unexpected saved tasks: %#v", savedTasks)
	}
}

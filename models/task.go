package models

import "time"

type Task struct {
	Name      string
	ID        int
	DueDate   time.Time
	Objective string
}

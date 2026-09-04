package models

import "time"

type Task struct {
	ID        int
	DueDate   time.Time
	Objective string
	Complete  bool
}

package models

import "time"

type Job struct {
	ID      string
	Name    string
	Comment string
	StartAt time.Time
	EndAt   time.Time
}

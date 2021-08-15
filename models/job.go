package models

import "time"

type Job struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Comment string    `json:"comment"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at,omitempty"`
}

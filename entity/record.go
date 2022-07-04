package entity

import "time"

type status string

const (
	StatusActive    status = "active"
	StatusCompleted status = "completed"
	StatusDeclined  status = "declined"
)

type Record struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Status      status    `json:"status"`
}

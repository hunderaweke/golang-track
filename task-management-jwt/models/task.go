package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	UserID      string    `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

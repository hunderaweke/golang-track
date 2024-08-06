package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	UserID      string    `bson:"user_id" json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

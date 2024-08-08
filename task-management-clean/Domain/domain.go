package domain

import (
	"context"
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	UserID      string    `bson:"user_id" json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsAdmin  bool   `bson:"is_admin" json:"-"`
}
type UserRepository interface {
	Create(c context.Context, u User) error
	Get(c context.Context) ([]User, error)
	GetByID(c context.Context) (User, error)
	Delete(c context.Context, userID string) error
	Update(c context.Context, data User) (User, error)
}
type TaskRepository interface {
	Create(c context.Context, t Task) error
	Get(c context.Context) ([]Task, error)
	GetByID(c context.Context, taskID string) (Task, error)
	GetByUserID(c context.Context, userID string) ([]Task, error)
	Update(c context.Context, data Task) (Task, error)
	Delete(c context.Context, taskID string) error
}

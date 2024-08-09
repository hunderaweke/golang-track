package domain

import (
	"context"
	"time"
)

const (
	UserCollection = "users"
	TaskCollection = "tasks"
)

type Task struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `bson:"user_id" json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}
type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsAdmin  bool   `bson:"is_admin" json:"-"`
}
type UserRepository interface {
	Create(c context.Context, u User) (*User, error)
	Get(c context.Context) ([]User, error)
	GetByID(c context.Context, userID string) (*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	PromoteUser(c context.Context, userID string) error
	Delete(c context.Context, userID string) error
	Update(c context.Context, userID string, data User) (*User, error)
}
type UserUsecase interface {
	Create(u User) (*User, error)
	Get() ([]User, error)
	GetByID(userID string) (*User, error)
	GetByEmail(email string) (*User, error)
	PromoteUser(userID string) error
	Delete(userID string) error
	Update(userID string, data User) (*User, error)
}
type TaskRepository interface {
	Create(c context.Context, t Task) (Task, error)
	Get(c context.Context) ([]Task, error)
	GetByID(c context.Context, taskID string) (Task, error)
	GetByUserID(c context.Context, userID string) ([]Task, error)
	Update(c context.Context, taskID string, data Task) (*Task, error)
	Delete(c context.Context, taskID string) error
}
type TaskUsecase interface {
	Create(t Task) (Task, error)
	Get() ([]Task, error)
	GetByID(taskID string) (Task, error)
	GetByUserID(userID string) ([]Task, error)
	Update(taskID string, data Task) (*Task, error)
	Delete(taskID string) error
}

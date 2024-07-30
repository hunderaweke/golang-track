package models

type Member struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	BorrowedBook []Book `json:"borrowed_book"`
}

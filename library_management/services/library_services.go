package services

import (
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Members map[int]models.Member
	Books   map[int]models.Book
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
}

func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}

func (l *Library) RemoveBook(book models.Book) {
	delete(l.Books, book.ID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	m, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("Member with id %d not found", memberID)
	}
	b, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("Book with id %d not found", bookID)
	}
	if b.Status == "Borrowed" {
		return fmt.Errorf("Book with %d is already borrowed", bookID)
	}
	m.BorrowedBook = append(m.BorrowedBook, b)
	l.Members[m.ID] = m
	b.Status = "Borrowed"
	l.Books[bookID] = b
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	m, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberID)
	}
	b, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id %d not found", bookID)
	}
	i := -1
	for j, book := range m.BorrowedBook {
		if book.ID == bookID {
			i = j
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("book with id %d was not borrowed by user with id %d", bookID, memberID)
	}
	m.BorrowedBook = append(m.BorrowedBook[:i], m.BorrowedBook[i+1:]...)
	b.Status = "Available"
	l.Books[bookID] = b
	l.Members[memberID] = m
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	a := make([]models.Book, 0)
	for _, b := range l.Books {
		if b.Status == "Available" {
			a = append(a, b)
		}
	}
	return a
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	return l.Members[memberID].BorrowedBook
}

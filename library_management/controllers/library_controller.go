package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strings"
)

var (
	L            services.Library = services.Library{Members: make(map[int]models.Member), Books: make(map[int]models.Book)}
	nextBookID   int              = 1
	nextMemberID int              = 1
)

func AddBook() {
	var title, author string
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter book title: ")
	title, _ = r.ReadString('\n')
	title = strings.Trim(title, "\n")
	fmt.Print("Enter book author: ")
	author, _ = r.ReadString('\n')
	author = strings.Trim(author, "\n")
	b := models.Book{ID: nextBookID, Title: title, Author: author}
	L.AddBook(b)
	fmt.Printf("Book added with ID %d\n", nextBookID)
	nextBookID++
}

func GetBooks() {
	fmt.Println("List of books:")
	for id, book := range L.Books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", id, book.Title, book.Author)
	}
}

func GetAvailableBooks() {
	fmt.Println("List of available books:")
	availableBooks := L.ListAvailableBooks()
	for _, book := range availableBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func GetBorrowedBooks() {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	borrowedBooks := L.ListBorrowedBooks(memberID)
	fmt.Printf("List of books borrowed by member %d:\n", memberID)
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func AddMember() {
	var name string
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter member name: ")
	name, _ = r.ReadString('\n')
	name = strings.Trim(name, "\n")
	m := models.Member{ID: nextMemberID, Name: name}
	L.AddMember(m)
	fmt.Printf("Member added with ID %d\n", nextMemberID)
	nextMemberID++
}

func GetMembers() {
	fmt.Println("List of members:")
	for id, member := range L.Members {
		fmt.Printf("ID: %d, Name: %s\n", id, member.Name)
	}
}

func BorrowBook() {
	var bookID, memberID int
	fmt.Print("Enter book ID: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	err := L.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func ReturnBook() {
	var bookID, memberID int
	fmt.Print("Enter book ID: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	err := L.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

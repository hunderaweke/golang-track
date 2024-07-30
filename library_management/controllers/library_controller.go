package main

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"strconv"
)

var (
	L            services.Library = services.Library{Members: make(map[int]models.Member), Books: make(map[int]models.Book)}
	nextBookID   int              = 1
	nextMemberID int              = 1
)

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Get Books")
		fmt.Println("3. Get Available Books")
		fmt.Println("4. Get Borrowed Books")
		fmt.Println("5. Add Member")
		fmt.Println("6. Get Members")
		fmt.Println("7. Borrow Book")
		fmt.Println("8. Return Book")
		fmt.Println("9. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook()
		case 2:
			getBooks()
		case 3:
			getAvailableBooks()
		case 4:
			getBorrowedBooks()
		case 5:
			addMember()
		case 6:
			getMembers()
		case 7:
			borrowBook()
		case 8:
			returnBook()
		case 9:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func addBook() {
	var title, author string
	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&author)

	b := models.Book{ID: nextBookID, Title: title, Author: author}
	L.AddBook(b)
	fmt.Printf("Book added with ID %d\n", nextBookID)
	nextBookID++
}

func getBooks() {
	fmt.Println("List of books:")
	for id, book := range L.Books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", id, book.Title, book.Author)
	}
}

func getAvailableBooks() {
	fmt.Println("List of available books:")
	availableBooks := L.ListAvailableBooks()
	for _, book := range availableBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func getBorrowedBooks() {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	borrowedBooks := L.ListBorrowedBooks(memberID)
	fmt.Printf("List of books borrowed by member %d:\n", memberID)
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func addMember() {
	var name string
	fmt.Print("Enter member name: ")
	fmt.Scanln(&name)

	m := models.Member{ID: nextMemberID, Name: name}
	L.AddMember(m)
	fmt.Printf("Member added with ID %d\n", nextMemberID)
	nextMemberID++
}

func getMembers() {
	fmt.Println("List of members:")
	for id, member := range L.Members {
		fmt.Printf("ID: %d, Name: %s\n", id, member.Name)
	}
}

func borrowBook() {
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

func returnBook() {
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

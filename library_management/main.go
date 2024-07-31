package main

import (
	"fmt"
	"library_management/controllers"
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
			controllers.AddBook()
		case 2:
			controllers.GetBooks()
		case 3:
			controllers.GetAvailableBooks()
		case 4:
			controllers.GetBorrowedBooks()
		case 5:
			controllers.AddMember()
		case 6:
			controllers.GetMembers()
		case 7:
			controllers.BorrowBook()
		case 8:
			controllers.ReturnBook()
		case 9:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

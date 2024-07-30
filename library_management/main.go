package main

import (
	"fmt"
	"library_management/controllers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /books/", controllers.AddBook)
	mux.HandleFunc("POST /members/{$}", controllers.AddMember)
	mux.HandleFunc("GET /books/{$}", controllers.GetBooks)
	mux.HandleFunc("GET /books/{id}", controllers.GetBooks)
	mux.HandleFunc("GET /members/{$}", controllers.GetMembers)
	mux.HandleFunc("GET /members/{id}", controllers.GetMembers)
	mux.HandleFunc("POST /members/{memberID}/borrow/{bookID}", controllers.BorrowBook)
	mux.HandleFunc("PUT /members/{memberID}/return/{bookID}", controllers.ReturnBook)
	mux.HandleFunc("GET /members/{memberID}/books/", controllers.GetBorrowedBooks)
	fmt.Println("Serving at host.... 9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}

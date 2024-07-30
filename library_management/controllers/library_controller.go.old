package controllers

import (
	"encoding/json"
	"fmt"
	"library_management/models"
	"library_management/services"
	"net/http"
	"strconv"
)

var (
	L            services.Library = services.Library{Members: make(map[int]models.Member), Books: make(map[int]models.Book)}
	nextBookID   int              = 1
	nextMemberID int              = 1
)

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %v", err)
	}
	return v, nil
}

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %v", err)
	}
	return nil
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	b, err := decode[models.Book](r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b.ID = nextBookID
	L.AddBook(b)
	encode(w, r, http.StatusCreated, L.Books[b.ID])
	nextBookID++
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		encode(w, r, http.StatusOK, L.Books)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b, ok := L.Books[intId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
	encode(w, r, http.StatusOK, b)
}

func GetAvailableBooks(w http.ResponseWriter, r *http.Request) {
	b := L.ListAvailableBooks()
	encode(w, r, http.StatusOK, b)
}

func GetBorrowedBooks(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("memberID")
	intID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	b := L.ListBorrowedBooks(intID)
	encode(w, r, http.StatusOK, b)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		encode(w, r, http.StatusOK, L.Members)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	m, ok := L.Members[intId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
	encode(w, r, http.StatusOK, m)
}

func AddMember(w http.ResponseWriter, r *http.Request) {
	m, err := decode[models.Member](r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	m.ID = nextMemberID
	L.AddMember(m)
	encode(w, r, http.StatusCreated, m)
	nextMemberID++
}

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PathValue("bookID")
	memberID := r.PathValue("memberID")
	if bookID == "" || memberID == "" {
		w.Write([]byte("404 Not Found"))
		return
	}
	intMemberID, err := strconv.Atoi(memberID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	intBookID, err := strconv.Atoi(bookID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, ok := L.Members[intMemberID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
	_, ok = L.Books[intBookID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
	if err = L.BorrowBook(intBookID, intMemberID); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	m, _ := L.Members[intMemberID]
	encode(w, r, http.StatusOK, m)
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PathValue("bookID")
	memberID := r.PathValue("memberID")
	if bookID == "" || memberID == "" {
		w.Write([]byte("404 Not Found"))
		return
	}
	intMemberID, err := strconv.Atoi(memberID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	intBookID, err := strconv.Atoi(bookID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, ok := L.Members[intMemberID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
	_, ok = L.Books[intBookID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
	if err = L.ReturnBook(intBookID, intMemberID); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	m, _ := L.Members[intMemberID]
	encode(w, r, http.StatusOK, m)
}

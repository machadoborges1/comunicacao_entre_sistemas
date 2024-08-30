package main

import (
	"encoding/json"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = map[string]Book{
	"123": {ID: "123", Title: "Go Programming", Author: "John Doe"},
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	book, ok := books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books[book.ID] = book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", getBookHandler).Methods("GET")
	r.HandleFunc("/books", createBookHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

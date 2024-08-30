package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Link struct {
	Href   string `json:"href"`
	Method string `json:"method"`
}

type Book struct {
	ID     string          `json:"id"`
	Title  string          `json:"title"`
	Author string          `json:"author"`
	Links  map[string]Link `json:"_links"`
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
	book.Links = map[string]Link{
		"self":   {Href: "/books/" + book.ID, Method: "GET"},
		"update": {Href: "/books/" + book.ID, Method: "PUT"},
		"delete": {Href: "/books/" + book.ID, Method: "DELETE"},
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", getBookHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

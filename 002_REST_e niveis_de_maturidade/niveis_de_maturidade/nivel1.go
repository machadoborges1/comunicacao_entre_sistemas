//Recursos: A API trata livros como recursos individuais acessíveis por URIs únicas (/books/{id}), mas não utiliza verbos HTTP padronizados para diferentes operações.

package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	id := r.URL.Path[len("/books/"):]
	book, ok := books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/books/", getBookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package handlers

import (
    "encoding/json"
    "go-library-api/data"
    "go-library-api/models"
    "net/http"
    "github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    book.ID = models.NewBook(book.Title, book.Author, book.PublishedYear).ID
    data.AddBook(book)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data.Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    book, err := data.GetBookByID(vars["id"])
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    err := data.DeleteBook(vars["id"])
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

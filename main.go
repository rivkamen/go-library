package main

import (
    "go-library-api/handlers"
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
    router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
    router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
    router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

    log.Println("Server running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

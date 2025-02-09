package models

import "github.com/google/uuid"

type Book struct {
    ID            string `json:"id"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedYear int    `json:"publishedYear"`
}

func NewBook(title, author string, year int) Book {
    return Book{
        ID:            uuid.New().String(),
        Title:         title,
        Author:        author,
        PublishedYear: year,
    }
}




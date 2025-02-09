package data

import (
    "go-library-api/models"
    "errors"
)

var Books = []models.Book{}

func AddBook(book models.Book) {
    Books = append(Books, book)
}

func GetBookByID(id string) (*models.Book, error) {
    for _, book := range Books {
        if book.ID == id {
            return &book, nil
        }
    }
    return nil, errors.New("book not found")
}

func DeleteBook(id string) error {
    for i, book := range Books {
        if book.ID == id {
            Books = append(Books[:i], Books[i+1:]...)
            return nil
        }
    }
    return errors.New("book not found")
}

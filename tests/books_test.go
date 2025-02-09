
package tests

import (
	"bytes"
	"encoding/json"
	"go-library-api/handlers"
	"go-library-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	book := models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2023,
	}

	body, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	handlers.CreateBook(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", res.Code)
	}
}

func TestGetBooks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/books", nil)
	res := httptest.NewRecorder()

	handlers.GetBooks(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.Code)
	}
}

func TestGetBookByID(t *testing.T) {
	book := models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2023,
	}
	handlers.CreateBook(httptest.NewRecorder(), createRequest(book))

	req, _ := http.NewRequest("GET", "/books/1", nil)
	res := httptest.NewRecorder()

	handlers.GetBookByID(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.Code)
	}
}

func TestDeleteBook(t *testing.T) {
	book := models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2023,
	}
	handlers.CreateBook(httptest.NewRecorder(), createRequest(book))

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	res := httptest.NewRecorder()

	handlers.DeleteBook(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", res.Code)
	}
}


func createRequest(book models.Book) *http.Request {
	body, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

# GoLang Developer Code Review

## GoLang Coding Assessment: REST API Implementation

### Objective
Build a simple REST API for a library system that manages books and authors.

## Features & Requirements

### API Endpoints

- **POST /books** - Add a new book. The request should include:

  ```json
  {
    "title": "Book Title",
    "author": "Author Name",
    "publishedYear": 2022
  }
  ```

- **GET /books** - Retrieve all books.
- **GET /books/{id}** - Retrieve a book by its ID.
- **DELETE /books/{id}** - Delete a book by its ID.

### Implementation Details

- Use **Gorilla Mux** (or any preferred router library) for routing.
- Store data in memory using a map or a slice.
- Implement basic error handling:
  - Return **404 Not Found** if the book does not exist.
  - Return appropriate status codes for all operations.

### Bonus Features

- Add **search functionality** (e.g., `/books?author=Author Name`).
- Use proper **Go practices**, such as structuring the project into packages.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation & Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/golang-library-api.git
   cd golang-library-api
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

4. The API should now be running at `http://localhost:8080`.

### Running Tests

Run unit tests using the built-in Go testing package:
```sh
  go test ./...
```

## Project Structure

```
├── main.go
├── handlers
│   ├── book_handler.go
├── models
│   ├── book.go
├── routes
│   ├── routes.go
├── storage
│   ├── memory_storage.go
├── tests
│   ├── book_handler_test.go
├── go.mod
└── go.sum
```

## Technologies Used
- **Go** - Programming language
- **Gorilla Mux** - Router library for handling endpoints
- **Testing package** - Built-in Go testing

## Future Enhancements
- Connect to a database (PostgreSQL, MongoDB, etc.) instead of in-memory storage.
- Implement authentication and authorization.
- Deploy to a cloud service.

## License
This project is open-source and available under the [MIT License](LICENSE).

---
### Author
[Your Name](https://github.com/yourusername)

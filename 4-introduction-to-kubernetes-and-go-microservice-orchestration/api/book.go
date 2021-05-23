package api

import (
	json "encoding/json"
	"io/ioutil"
	http "net/http"
)

// Book type with Name, Author and ISBN
type Book struct {
	// define the book
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

const (
	content_type_json string = "application/json; charset=utf-8"
)

// toJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJson
}

// FromJson to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = map[string]Book{
	"034567802": {Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "034567802"},
	"000000000": {Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "000000000"},
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// BookHandleFunc to be used as http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// AllBooks returns slice of all books
func AllBooks() []Book {
	books := []Book{}
	for _, book := range Books {
		books = append(books, book)
	}
	return books
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	if book, exists := Books[isbn]; exists {
		return book, true
	} else {
		return Book{}, false
	}
}

// CreateBook add a book to the books map
func CreateBook(b Book) (isbn string, ok bool) {
	isbn = b.ISBN
	if _, ok := Books[isbn]; ok {
		return "", false
	}

	Books[isbn] = b
	return isbn, true
}

// DeleteBook add a book to the books map
func DeleteBook(isbn string) {
	delete(Books, isbn)
}

// UpdateBook add a book to the books map
func UpdateBook(isbn string, b Book) (exists bool) {
	if _, exists := Books[isbn]; exists {
		Books[isbn] = b
	}

	return exists
}

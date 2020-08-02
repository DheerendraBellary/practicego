package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

var books map[string]Book

func ToJSON(b Book) ([]byte, error) {
	return json.Marshal(b)
}

func FromJSON(bJson []byte) (*Book, error) {
	book := &Book{}
	err := json.Unmarshal(bJson, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func AddBook(b *Book) error {
	if _, ok := books[b.ISBN]; ok {
		return errors.New("Book with ISBN already exists")
	}
	books[b.ISBN] = *b
	return nil
}

func AllBooks() []Book {
	ret := []Book{}
	for _, book := range books {
		ret = append(ret, book)
	}
	return ret
}

func UpdateBook(isbn string, book *Book) {
	book.ISBN = isbn
	books[isbn] = *book
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		booksJson, _ := json.Marshal(books)
		w.Header().Add("content-type", "application/json")
		w.Header().Add("charset", "utf-8")
		w.Write(booksJson)

	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		b, err := FromJSON(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = AddBook(b)
		if err != nil {
			fmt.Printf("BooksHandler: Method: http.POST. isbn: %v already exists\n", b.ISBN)
			w.WriteHeader(http.StatusConflict)
		}

		w.Header().Add("Location", "/api/books/"+b.ISBN)
		w.WriteHeader(http.StatusCreated)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/book/"):]

	switch method := r.Method; method {

	case http.MethodGet:
		if book, ok := books[isbn]; ok {
			bookJson, _ := json.Marshal(book)
			w.Header().Add("content-type", "application/json")
			w.Header().Add("charset", "utf-8")
			w.Write(bookJson)
			return
		}
		w.WriteHeader(http.StatusNotFound)

	case http.MethodPut:
		if _, ok := books[isbn]; ok {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			b, err := FromJSON(body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			UpdateBook(isbn, b)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)

	case http.MethodDelete:
		if _, ok := books[isbn]; ok {
			delete(books, isbn)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return

}

func init() {
	books = map[string]Book{
		"100001": {"Basics of C++", "Balaguruswamy", "100001"},
		"100002": {"Data Structure and Algorithms", "Kosaraju", "100002"},
	}
}

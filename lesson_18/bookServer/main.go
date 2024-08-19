package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    int      `json:"year"`
}

var Books = map[int]Book{
	1: Book{
		Id:      1,
		Title:   "Go на практике",
		Authors: []string{"Мэтт Батчер", "Мэтт Фарина"},
		Year:    2016,
	},
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	var newBook Book
	err = json.Unmarshal(jsong, &newBook)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}
	t := newBook.Title
	a := newBook.Authors
	y := newBook.Year
	if t == "" {
		handleError(w, http.StatusBadRequest, fmt.Errorf("title not found"))
		return
	}
	if a == nil {
		handleError(w, http.StatusBadRequest, fmt.Errorf("authors not found"))
		return
	}
	if y == 0 {
		handleError(w, http.StatusBadRequest, fmt.Errorf("year not found"))
		return
	}

	newBook.Id = len(Books) + 1
	Books[len(Books)+1] = newBook
	data, err := json.Marshal(newBook)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}
	book, ok := Books[idint]
	if !ok {
		handleError(w, http.StatusNotFound, fmt.Errorf("id %d not found", idint))
		return
	}
	data, err := json.Marshal(book)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}
	delete(Books, idint)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	var book Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}
	Books[book.Id] = book
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	limit := query.Get("limit")

	limitBooks := make(map[int]Book)
	if limit != "" {
		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			handleError(w, http.StatusBadRequest, errors.New("invalid limit parameter"))
			return
		}
		if limitNum > len(Books) {
			limitNum = len(Books)
		}
		for i := 1; i < limitNum+1; i++ {
			limitBooks[i] = Books[i]
		}
	} else {
		limitBooks = Books
	}

	data, err := json.Marshal(limitBooks)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book", GetBook).Methods(http.MethodGet)
	r.HandleFunc("/book", AddBook).Methods(http.MethodPost)
	r.HandleFunc("/book", DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/book", UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books", AllBooks).Methods(http.MethodGet)
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func handleError(w http.ResponseWriter, status int, err error) {
	result := map[string]interface{}{
		"error":  err.Error(),
		"status": http.StatusText(status),
	}
	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

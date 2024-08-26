package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/internal/db"
	"server/internal/domain"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	var newBook domain.Book
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

	newBook.Id = len(db.Books) + 1
	db.Books[len(db.Books)+1] = newBook
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
	book, ok := db.Books[idint]
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
	delete(db.Books, idint)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	var book domain.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}
	db.Books[book.Id] = book
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	limit := query.Get("limit")

	limitBooks := make(map[int]domain.Book)
	if limit != "" {
		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			handleError(w, http.StatusBadRequest, errors.New("invalid limit parameter"))
			return
		}
		if limitNum > len(db.Books) {
			limitNum = len(db.Books)
		}
		for i := 1; i < limitNum+1; i++ {
			limitBooks[i] = db.Books[i]
		}
	} else {
		limitBooks = db.Books
	}

	data, err := json.Marshal(limitBooks)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)
}

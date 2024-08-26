package db

import "server/internal/domain"

var Books = map[int]domain.Book{
	1: {
		Id:      1,
		Title:   "Go на практике",
		Authors: []string{"Мэтт Батчер", "Мэтт Фарина"},
		Year:    2016,
	},
}

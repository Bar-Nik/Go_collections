package main

import "fmt"

// 1. Создание и использование структур: Создайте структуру,
// представляющую собой "Book" с полями "Title", "Author" и "Pages".
// Создайте экземпляр этой структуры, заполните поля и выведите их на
// экран.

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	b := Book{
		Title:  "Дом Драконов",
		Author: "Мартин Джордж",
		Pages:  768,
	}
	fmt.Print("Название книги: ", b.Title, "\n", "Автор книги: ", b.Author, "\n", "Кол-во страниц: ", b.Pages, "\n")

}

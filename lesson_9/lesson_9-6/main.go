package main

import "fmt"

// 6. Структуры и функции (сложное) Создайте функцию, которая
// 	принимает два аргумента типа "Book" и возвращает книгу с большим
// 	количеством страниц. Протестируйте эту функцию с несколькими
// 	экземплярами структуры "Book".

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {

	b1 := Book{
		Title:  "Дом Драконов",
		Author: "Мартин Джордж",
		Pages:  768,
	}

	b2 := Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Pages:  1300,
	}
	fmt.Println(PageComparison(b1, b2).Title)
}

func PageComparison(a Book, b Book) Book {
	if a.Pages > b.Pages {
		return a
	} else {
		return b
	}
}

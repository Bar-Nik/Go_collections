package main

import "fmt"

// Добавление элементов в слайс: Создайте слайс, добавьте в него
// несколько элементов с помощью функции append и выведите все
// элементы на экран.

func main() {
	sl := []int{443, 876}
	sl = append(sl, 765, 9087)
	fmt.Println(sl)
}

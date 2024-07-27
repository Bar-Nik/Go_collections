package main

import "fmt"

// Копирование слайсов: Создайте два слайса, скопируйте элементы из
// одного слайса в другой с помощью функции copy и выведите оба слайса
// на экран.

func main() {
	sl_one := []int{1, 2, 3, 5, 8}
	sl_two := make([]int, len(sl_one))
	copy(sl_two, sl_one)
	fmt.Println("Первый слайс:\n", sl_one)
	fmt.Println("Второй слайс:\n", sl_two)
}

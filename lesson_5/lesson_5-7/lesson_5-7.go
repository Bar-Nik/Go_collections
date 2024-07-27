package main

import "fmt"

// 7. Цикл for-range с слайсом: Создайте слайс с несколькими элементами.
// Напишите цикл for-range, который выводит каждый элемент слайса на
// экран.

func main() {
	sl := []int{1, 3, 5, 7}
	for _, v := range sl {
		fmt.Println(v)
	}
}

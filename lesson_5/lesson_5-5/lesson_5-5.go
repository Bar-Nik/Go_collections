package main

import "fmt"

// 5. Цикл for-range с массивом: Создайте массив с несколькими элементами.
// Напишите цикл for-range, который выводит каждый элемент массива на
// экран.

func main() {
	array := [4]int{100, 200, 300, 400}
	for v := range array {
		fmt.Println(v)
	}
}

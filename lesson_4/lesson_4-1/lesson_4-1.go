package main

import "fmt"

// 1. Объявление массивов: Создайте массив из пяти целых чисел и
// присвойте каждому элементу значение. Выведите каждый элемент на
// экран отдельно.

func main() {
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(numbers[3])
	fmt.Println(numbers[4])
}

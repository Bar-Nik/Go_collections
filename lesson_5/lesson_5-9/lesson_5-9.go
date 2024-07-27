package main

import "fmt"

// 9. Наполнение массива: Создайте массив типа int с 5 элементами, с
// помощью цикла запишите в массив значения, которые будут равняться
// индексу, по которому значение будет находиться

func main() {

	array := [5]int{}
	fmt.Println(array)
	for i := range array {
		array[i] = i
	}
	fmt.Println(array)
}

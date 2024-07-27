package main

import "fmt"

// 8. Изменение нескольких элементов массивы: Создайте массив типа int с
// элементами 1, 2, 3, 4, измените их порядок в массиве на 4, 3, 2, 1 с
// помощью изменения элементов по индексу.

func main() {
	array := [4]int{1, 2, 3, 4}
	fmt.Println(array)
	array[0] = 4
	array[1] = 3
	array[2] = 2
	array[3] = 1
	fmt.Println(array)

}

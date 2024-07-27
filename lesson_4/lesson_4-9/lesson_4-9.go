package main

import "fmt"

// Сохранение значения переменной в массив: Создайте 3 переменных с
// разными значениями типа string, создайте массив типа string с 3
// элементами, положите в массив значения переменных под разные
// индексы.

func main() {
	str_1 := "String-1"
	str_2 := "String-2"
	str_3 := "String-3"
	var array [3]string
	array[0] = str_2
	array[1] = str_1
	array[2] = str_3
	fmt.Println(array)

}

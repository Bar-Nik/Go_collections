package main

import "fmt"

// 1. Создание слайсов: Объявите слайс целых чисел, добавьте в него
// несколько значений и выведите эти значения на экран.

func main() {
	sli := []int{33, 22}
	fmt.Println(sli)
	sli = append(sli, 44, 88)
	fmt.Println(sli)

}

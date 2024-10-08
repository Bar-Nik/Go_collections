package main

import "fmt"

// 5. Оператор выбора switch с несколькими значениями для case: Создайте
// слайс из чисел. Напишите цикл, который использует оператор switch для
// проверки каждого числа и выводит сообщение для четных и нечетных
// чисел.

func main() {
	slice := []int{11, 22, 33, 44, 55}
	for _, k := range slice {
		switch k % 2 {
		case 0:
			fmt.Println(k, "- четное")
		case 1:
			fmt.Println(k, "- нечетное")
		}
	}
}

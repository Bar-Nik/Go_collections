package main

import "fmt"

// 7. Оператор выбора switch с fallthrough (сложное) Создайте слайс из
// 	чисел от 1 до 10. Напишите цикл, который использует оператор switch с
// 	fallthrough для проверки каждого числа и выводит сообщение для
// 	каждого случая.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, k := range slice {
		switch k {
		case 1:
			fmt.Println("Первый")
		case 2:
			fmt.Println("Второй")
		case 3:
			fmt.Println("Третий")
			fallthrough
		case 4:
			fmt.Println("Четвертый")
		case 5:
			fmt.Println("Пятый")
		case 6:
			fmt.Println("Шестой")
		case 7:
			fmt.Println("Седьмой")
			fallthrough
		case 8:
			fmt.Println("Восьмой")
		case 9:
			fmt.Println("Девятый")
		case 10:
			fmt.Println("Десятый")
		}
	}
}

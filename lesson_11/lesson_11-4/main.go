package main

import "fmt"

// 4. Оператор выбора switch в цикле: Создайте слайс из чисел от 1 до 7.
// Напишите цикл, который использует оператор switch для проверки
// каждого числа как дня недели и выводит сообщение в зависимости от
// дня.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	for _, i := range slice {
		switch i {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		}
	}
}

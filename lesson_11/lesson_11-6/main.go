package main

import "fmt"

// 6. Оператор выбора switch без выражения: Создайте слайс из булевых
// значений. Напишите цикл, который использует оператор switch без
// выражения для вывода разных сообщений в зависимости от значения
// булевой переменной.

func main() {
	slice := []bool{true, false, false, true, true}
	for _, k := range slice {
		switch k {
		case k:
			fmt.Println("Значение верно")
		default:
			fmt.Println("Значение неверно")
		}
	}
}

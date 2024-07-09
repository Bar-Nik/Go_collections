package main

import "fmt"

// 2. Функция с параметрами: Создайте функцию, которая принимает два
// числа в качестве аргументов и выводит их сумму.

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Println("Сумма:", sum(a, b))
}
func sum(t int, d int) int {
	return t + d

}

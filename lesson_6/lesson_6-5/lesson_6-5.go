package main

import "fmt"

// 5.Функция с переменным числом аргументов: Создайте функцию, которая
// принимает переменное количество аргументов и выводит их на экран.

func main() {
	func1("One", "Two")
}
func func1(numes ...string) {
	for _, num := range numes {
		fmt.Println(num)
	}

}

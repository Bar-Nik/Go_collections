package main

import "fmt"

// 7. Функция, возвращающая сумму и количество аргументов: Создайте
// функцию, которая получает переменное количество аргументов типа int,
// складывает их, и возвращает их сумму и количество полученных ею
// аргументов.

func main() {
	a, b := SumCount(1, 2, 3, 4)
	fmt.Println("Сумма чисел:", a, "Количество чисел", b)
}

func SumCount(nums ...int) (int, int) {
	var sum, count int
	for _, num := range nums {
		sum += num
		count++
	}
	return sum, count
}

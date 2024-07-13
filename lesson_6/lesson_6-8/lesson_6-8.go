package main

import "fmt"

// 8. Возвращение большего значения из функции: Создайте функцию,
// которая принимает 3 переменных типа int, значение передаваемое в
// функцию должно быть 0 или 1, и возвращает значение, которое
// повторяется большее количество раз, среди полученных переменных.

func main() {
	var a, b, c int
	fmt.Println("Введите числа 0 или 1 три раза")
	fmt.Scan(&a, &b, &c)
	fmt.Println(ReturnInt(a, b, c))

}

func ReturnInt(nums ...int) int {
	var one, zero int
	for num := range nums {
		if num != 0 {
			one += 1
		} else {
			zero += 1
		}
	}

	if one > zero {
		return 1
	} else {
		return 0
	}

}

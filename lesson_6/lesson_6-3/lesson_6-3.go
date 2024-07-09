package main

import "fmt"

// 3. Функция, возвращающая значение: Создайте функцию, которая
// принимает два числа и возвращает их произведение.

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Println("Произведение двух значений:", mult(a, b))

}
func mult(r int, d int) int {
	return r * d
}

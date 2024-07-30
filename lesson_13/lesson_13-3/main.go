package main

import "fmt"

// 3. Возвращение указателя из функции: Напишите функцию, которая
// возвращает указатель на новую переменную int.

func MyFunc(x int) *int {
	l := new(int)
	*l = x
	return l

}

func main() {
	x := MyFunc(7)
	fmt.Printf("Значение переменной: %d\nАдрес переменной: %x\n", *x, x)

}

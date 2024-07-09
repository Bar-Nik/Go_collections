package main

import "fmt"

// 6. Вызов функции из другой функции: Создайте две функции и вызовите
// одну из них из другой.

func main() {
	call_func1()
}

func call_func1() {
	fmt.Println("Функция call_func1")
	call_func2()
}

func call_func2() {
	fmt.Println("Функция call_func2")
}

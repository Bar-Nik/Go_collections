package main

import (
	"errors"
	"fmt"
	"log"
)

// 3. Проверка на наличие ошибки: Напишите функцию, которая возвращает
// ошибку в некоторых случаях. Вызовите эту функцию и проверьте, есть ли
// ошибка, прежде чем продолжать выполнение кода.

func main() {
	num, err := myfunc()
	if err != nil {
		fmt.Println("Число не подходит!")
		log.Fatal(err)
	}
	fmt.Println("Наше число:", num)
}

func myfunc() (int, error) {
	var a int
	var ErrInvalidInput = errors.New("invalid input")
	fmt.Scan(&a)
	if a < 0 {
		return a, ErrInvalidInput
	}
	return a, nil
}

package main

import (
	"fmt"
	"os"
	"time"
)

// 6. Создание кастомного типа ошибки ( сложное): Создайте свой
// собственный тип ошибки, реализовав интерфейс error. Создайте функцию,
// которая возвращает вашу кастомную ошибку, и обработайте её.

type MyError struct {
	Msg      string
	TimeHour int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s -- %d", e.Msg, e.TimeHour)
}

func main() {
	e := MyError{
		Msg:      "моя ошибка в",
		TimeHour: time.Now().Hour(),
	}
	x, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Сработала в ", e.Error())
	} else {
		fmt.Println(x)
	}

}

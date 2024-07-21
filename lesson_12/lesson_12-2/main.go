package main

import (
	"errors"
	"fmt"
)

// 2. Создание новой ошибки: Используйте функцию errors.New для создания
// новой ошибки. Верните эту ошибку из вашей функции и обработайте её.

func main() {
	e := errMy()
	if e != nil {
		fmt.Println(e)
	}

}
func errMy() error {
	var ErrInvalid = errors.New("Моя ошибка")
	return ErrInvalid
}

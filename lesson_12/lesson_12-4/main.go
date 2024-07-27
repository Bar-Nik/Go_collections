package main

import (
	"fmt"
	"log"
)

// 4. Использование fmt.Errorf для создания ошибки: Используйте функцию
// fmt.Errorf для создания новой ошибки с форматированным сообщением.
// Верните эту ошибку из вашей функции и обработайте её.

func main() {
	s, e := myFunc(2, 2)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(s)
	l, er := myFunc(10, 2)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(l)

}

func myFunc(x, y int) (int, error) {
	switch x {
	case 10:
		return 0, fmt.Errorf("число: %d не должно быть  равно 10", x)
	case 100:
		return 0, fmt.Errorf("число: %d не должно быть  равно 100", x)
	}
	return x * y, nil
}

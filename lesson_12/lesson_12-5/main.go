package main

import (
	"fmt"
	"log"
	"os"
)

// 5. Обработка ошибки из стандартной библиотеки: Вызовите функцию из
// стандартной библиотеки Go, которая может вернуть ошибку (например,
// os.Open). Обработайте возвращаемую ошибку.

func main() {
	x, e := os.Open("file.txt")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(x)
}

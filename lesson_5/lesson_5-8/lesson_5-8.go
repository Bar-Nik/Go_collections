package main

import "fmt"

// 8. Цикл for-range с строкой: Создайте строку. Напишите цикл for-range,
// который выводит каждый символ строки на экран.

func main() {
	str := "Golang"
	for _, v := range str {
		fmt.Println(string(v))
	}
}

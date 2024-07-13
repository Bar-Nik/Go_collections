package main

import "fmt"

// 6. Цикл for-range с map: Создайте map с несколькими парами ключ-
// значение. Напишите цикл for-range, который выводит каждую пару ключ-
// значение на экран.

func main() {
	dict := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	for k, v := range dict {
		fmt.Println("key:", k, "value:", v)
	}
}

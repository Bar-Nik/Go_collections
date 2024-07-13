package main

import "fmt"

// 2. Цикл for с условием: Напишите цикл for, который выводит на экран все
// четные числа от 1 до 20.

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

package main

import "fmt"

// 4. Цикл for с break: Напишите цикл for, который выводит числа от 1 до 100, но
// прерывает цикл, как только сумма выводимых чисел превысит 50.

func main() {
	count := 0
	for i := 0; i <= 100; i++ {
		count += i
		if count > 50 {
			break
		}
		fmt.Println(i)
	}
}

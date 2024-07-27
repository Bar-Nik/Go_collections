package main

import "fmt"

// 10. Цикл в цикле: Используя два цикла выведите таблицу умножения.

func main() {
	for i := 1; i <= 9; i++ {
		for t := 1; t <= 9; t++ {
			fmt.Print(i, "*", t, "=", i*t, ";  ")
		}
		fmt.Println()
	}
}

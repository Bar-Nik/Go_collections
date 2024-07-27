package main

import "fmt"

// 1. Создание map: Создайте map, где ключом является строка, а значением -
// целое число. Добавьте в нее несколько пар ключ-значение и выведите их
// на экран.

func main() {
	new_map := map[string]int{}
	new_map["Строка1"] = 1
	new_map["Строка2"] = 2
	fmt.Println(new_map)
}

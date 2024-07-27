package main

import "fmt"

// 8. Слайс как значение: Создайте map, где значением будут слайс типа int.
// Создайте несколько пар ключ-значение в данной map. Выведите
// результат.

func main() {
	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 5, 6}
	map_sl := map[string][]int{
		"one": sl1,
		"two": sl2,
	}
	fmt.Println(map_sl)

}

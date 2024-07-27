package main

import "fmt"

// 6. Срезы слайса: Создайте слайс с несколькими элементами, затем
// создайте новый слайс как срез первого слайса и выведите его на экран.

func main() {
	sl_one := []int{1, 2, 3, 5, 8}
	sl_two := sl_one[1:(len(sl_one) - 1)]
	fmt.Println(sl_one)
	fmt.Println(sl_two)
}

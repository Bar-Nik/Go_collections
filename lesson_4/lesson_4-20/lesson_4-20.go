package main

import "fmt"

// Минимальная длина и вместимость слайса: Создайте слайс с
// минимально возможной длиной и вместимостью, выведите слайс, его
// длину и вместимость.

func main() {
	sl := []int{}
	fmt.Println("Слайс:", sl, "\nДлина слайса:", len(sl), "\nВместимость слайса:", cap(sl))
}

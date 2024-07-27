package main

import "fmt"

// 2. Изменение элементов слайса: Создайте слайс с несколькими строками,
// измените одно из значений и выведите все элементы на экран.

func main() {
	sli := []string{"cat", "dog"}
	fmt.Println(sli)
	sli = append(sli, "elephant")
	fmt.Println(sli)
}

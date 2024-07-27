package main

import "fmt"

// Создание слайса с помощью make: Создайте слайс с помощью функции
// make, добавьте в него несколько элементов и выведите эти элементы на
// экран.

func main() {
	var sli []int
	sli = make([]int, 5, 23)
	sli = append(sli, 908799)
	sli = append(sli, 99987885)
	fmt.Println(sli)
}

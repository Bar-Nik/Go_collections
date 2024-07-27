package main

import "fmt"

// 3. Длина и вместимость слайса: Создайте слайс и добавьте в него
// несколько элементов. Используйте функции len и cap для определения
// длины и вместимости слайса. Выведите эти значения на экран.

func main() {
	sli := []int{}
	sli = append(sli, 1, 3, 77)
	fmt.Println("Длина слайса:", len(sli))
	fmt.Println("Вместимость слайса:", cap(sli))
	sli = append(sli, 9099)
	fmt.Println("Длина слайса:", len(sli))
	fmt.Println("Вместимость слайса:", cap(sli))

}

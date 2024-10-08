package main

import "fmt"

// Изменение вместимости слайса: Создайте слайс произвольного типа с 6
// элементами, узнайте его вместимость с помощью функции cap, опытным
// путём определите в каком случае вместимость слайса будет 12, 24 и
// более.

func main() {
	sl1 := []int{}
	fmt.Println("\nВместимость", cap(sl1))
	fmt.Println("Длина слайса", len(sl1))
	sl := []int{1, 22, 333, 4444, 55555, 666666}
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 7777777)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 88888888)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 999999999)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 1000000000)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 11111111111)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 222222222222)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 3333333333333)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 44444444444444)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))
	sl = append(sl, 99, 88, 77, 66, 55, 44, 33, 222, 111, 999, 888)
	fmt.Println("\nВместимость", cap(sl))
	fmt.Println("Длина слайса", len(sl))

}

// У пустого слайса вместимость и длина 0, после добавления элементов (увеличения длины слайса) в промежутке
// 0 до 6 включительно вместимость - 6, после добавления элементов (увеличения длины слайса) в промежутке от 7 до 12
// включительно вместимость - 12, после добавления элементов (увеличения длины слайса) в промежутке
// 13 до 24 включительно вместимость - 24 и т.д

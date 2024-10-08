package main

import "fmt"

// 11. Взаимосвязь слайсов и массивов ( сложное): Создайте массив типа int
// с 5 элементами 1, 2 3, 4, 5, после чего создайте слайс на основе первых
// трёх элементов созданного массива с помощью оператора слайса,
// добавьте к нему элемент 9 с помощью функции append(), после чего
// создайте новый слайс на основе первого и добавьте к нему два элемента
// 8 и 7, выведите массив и оба слайса на экран. Попробуйте объяснить
// полученный результат.

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Исходный массив:", array)
	sl := []int{array[0], array[1], array[2]}
	fmt.Println("Слайс на основе трех первых элементов массива:", sl)
	sl = append(sl, 9)
	fmt.Println("Слайс с добавлением нового элемента:", sl)
	sl_new := make([]int, len(sl))
	copy(sl_new, sl)
	fmt.Println("Новый слайс (копия):", sl_new)
	sl_new = append(sl_new, 8, 7)
	fmt.Println("Новый слайс с добавлением новых элементов:", sl_new)
}

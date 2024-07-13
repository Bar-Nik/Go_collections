package main

import "fmt"

// 4. Использование нескольких значений return: Создайте функцию, которая
// принимает два числа и возвращает их сумму и разность как два
// отдельных значения.

func main() {
	arr := nums(200, 400)
	fmt.Println("Сумма:", arr[0], "Разность:", arr[1])
	a, b := nums2(200, 400)
	fmt.Println("Сумма:", a, "Разность:", b)

}
func nums(v int, f int) [2]int {
	array := [2]int{}
	array[0] = v + f
	array[1] = v - f

	return array

}
func nums2(v int, f int) (int, int) {
	return v + f, v - f
}

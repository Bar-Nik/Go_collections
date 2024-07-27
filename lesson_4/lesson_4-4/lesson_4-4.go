package main

import "fmt"

// 4. Массивы с разными типами данных: Создайте три массива: один с
// целыми числами, один с числами с плавающей точкой и один со
// строками. Выведите все значения всех массивов.

func main() {
	array_int := [4]int{22, 44, 66, 77}
	array_float := [3]float64{12.1, 77.7, 12.3}
	array_str := [2]string{"milky", "way"}
	fmt.Println(array_int, array_float, array_str)
}

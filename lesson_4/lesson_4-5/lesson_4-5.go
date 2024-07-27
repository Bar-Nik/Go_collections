package main

import (
	"fmt"
)

// Многомерные массивы: Создайте двумерный массив (массив массивов)
// и присвойте каждому элементу значение. Выведите все значения на
// экран.

func main() {
	array := [4][2]int{}
	fmt.Println(array)
	array[0][0] = 1
	array[0][1] = 2
	array[1][0] = 3
	array[1][1] = 4
	array[2][0] = 4
	array[2][1] = 4
	array[3][0] = 4
	array[3][1] = 4

	fmt.Println(array)

}

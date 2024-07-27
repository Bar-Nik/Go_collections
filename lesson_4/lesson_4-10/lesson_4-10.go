package main

import (
	"fmt"
	"strconv"
)

// 10. Сохранение элемента массивы в переменную: Создайте массив типа int
// с 3 элементами, также создайте 3 переменные разных типов (int, string,
// float64). Сохраните элементы массива в переменные с помощью
// конвертации. Выведите результат.

func main() {
	array := [3]int{1, 2, 3}
	fmt.Println(array)
	var i int
	var s string
	var f float64
	i = array[0]
	s = strconv.Itoa(array[1])
	f = float64(array[2])
	fmt.Printf("int: %d, string: %s, float: %f", i, s, f)
}

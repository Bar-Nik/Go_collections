package main

import "fmt"

// 2. Изменение элементов массива: Создайте массив из трех строк.
// Измените значение второго элемента и выведите все элементы на экран.

func main() {
	array := [3]string{"just", "don't do", "it"}
	fmt.Println(array)
	array[1] = "DO"
	fmt.Println(array)
}

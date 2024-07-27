package main

import "fmt"

// 4. Удаление значения из map: Создайте map, добавьте в нее несколько пар
// ключ-значение. Затем удалите одну из пар с помощью функции delete и
// выведите всю map на экран.

func main() {
	capitals := map[string]string{
		"Russia":  "Moscow",
		"England": "London",
		"Spain":   "Madrid",
	}

	delete(capitals, "Russia")

	fmt.Println(capitals)
}

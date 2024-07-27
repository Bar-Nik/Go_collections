package main

import "fmt"

// 5. Проверка существования ключа в map: Создайте map, добавьте в нее
// несколько пар ключ-значение. Затем проверьте существование одного из
// ключей с помощью двухзначной формы индексирования и выведите
// результат на экран.

func main() {
	map_1 := map[string]string{
		"String-1": "Строка-1",
		"String-2": "Строка-2",
		"String-3": "Строка-3",
	}
	x, exist := map_1["String-1"]
	if exist {
		fmt.Println(x)
	} else {
		fmt.Println("Такого ключа нет!")
	}
}

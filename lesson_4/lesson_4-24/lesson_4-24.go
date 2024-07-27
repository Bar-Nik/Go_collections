package main

import "fmt"

// 3. Получение значения по ключу из map: Создайте map, добавьте в нее
// несколько пар ключ-значение. Затем получите и выведите на экран
// значение по одному из ключей.

func main() {
	capitals := map[string]string{
		"Russia": "Moscow",
	}
	capitals["England"] = "London"
	capitals["Spain"] = "Madrid"
	fmt.Println(capitals)
	fmt.Println(capitals["Russia"])
	fmt.Println(capitals["England"])
	fmt.Println(capitals["Spain"])

}

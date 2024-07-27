package main

import "fmt"

// 6. Проход по всем значениям map: Создайте map и добавьте в нее
// несколько пар ключ-значение. Затем напечатайте все пары ключ-
// значение. Для этого вам потребуется использовать цикл range (я знаю,
// что вы еще не изучали циклы, но это хорошая возможность начать).

func main() {
	capitals := map[string]string{
		"Russia":  "Moscow",
		"England": "London",
		"Spain":   "Madrid",
		"Italia":  "Rome",
		"France":  "Paris",
	}
	for country, capital := range capitals {
		fmt.Printf("Страна - %s, столица - %s\n", country, capital)
	}
}

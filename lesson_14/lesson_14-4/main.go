package main

import (
	"encoding/json"
	"fmt"
)

// 4. Обработка ошибок сериализации: Напишите функцию, которая
// принимает структуру и возвращает JSON-строку. Обработайте
// возможные ошибки при сериализации и выведите их на экран.

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {
	b := Book{ID: 1, Name: "Капитанская дочка", Author: "Пушкин"}
	fmt.Println(myFuncJson(b))

}
func myFuncJson(b Book) string {
	bookJSON, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return "Ошибка сериализации"
	}
	return string(bookJSON)
}

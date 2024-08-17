package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 5. Обработка ошибок десериализации: Напишите функцию, которая
// принимает JSON-строку и возвращает структуру. Обработайте
// возможные ошибки при десериализации и выведите их на экран.

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {
	b := Book{ID: 1, Name: "Капитанская дочка", Author: "Пушкин"}
	bJson := mySeriJson(b)

	fmt.Println(bJson)
	fmt.Println(myDESeriJson(bJson))

}
func mySeriJson(b Book) string {
	bookJSON, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return "Ошибка"
	}
	return string(bookJSON)
}

func myDESeriJson(b string) Book {
	var desb Book
	err := json.Unmarshal([]byte(b), &desb)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		log.Fatal(err)
	}
	return desb

}

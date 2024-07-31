package main

import (
	"encoding/json"
	"fmt"
)

// 1. Сериализация структуры: Создайте структуру Product с полями Name ,
// Price
// и Quantity . Создайте экземпляр этой структуры и сериализуйте его в
// JSON.

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {
	prod := Product{Name: "apple", Price: 12, Quantity: 2}
	prodJSON, err := json.Marshal(prod)
	if err != nil {
		fmt.Println("Ошибка сериализации", err)
		return
	}
	fmt.Println("JSON:", string(prodJSON))

}

package main

import (
	"encoding/json"
	"fmt"
)

// 2. Десериализация JSON: Создайте JSON-строку, представляющую данные
// о продукте. Десериализуйте эту строку в структуру Product и выведите
// поля структуры на экран.

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {
	prod := Product{Name: "jam", Price: 60, Quantity: 1}
	prodJSON, err := json.Marshal(prod)
	if err != nil {
		fmt.Println("Ошибка сериализации", err)
		return
	}
	fmt.Println("JSON:", string(prodJSON))
	var deccodProd Product
	err = json.Unmarshal(prodJSON, &deccodProd)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Println("Десериализованная структура:", deccodProd)

}

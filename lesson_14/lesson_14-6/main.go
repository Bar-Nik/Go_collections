package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 6. Сложная структура (⭐ сложное): Создайте структуру Order , которая
// 	содержит поле Products , являющееся срезом структур Product .
// 	Сериализуйте и десериализуйте экземпляр структуры Order .

type Product struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}
type Order struct {
	Products []Product `json:"products"`
}

func main() {
	a := Product{Name: "Cup", Amount: 4}
	b := Product{Name: "Pen", Amount: 10}

	products := Order{Products: []Product{a, b}}
	fmt.Println("Вывод структуры Order:", products)
	productsJson := mySeri(products)
	fmt.Println("JSON:", productsJson)
	productsOrder := myDESeri(productsJson)
	fmt.Println("Вывод структуры Order:", productsOrder)
}

func mySeri(p Order) string {
	bJson, err := json.Marshal(p)
	if err != nil {
		return "Ошибка сериализации"
	}
	return string(bJson)
}

func myDESeri(s string) Order {
	var d Order
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		log.Fatal(err)
	}
	return d

}

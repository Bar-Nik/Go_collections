package main

import "fmt"

// 4. Создание структуры с вложенными структурами (⭐ сложное):
// 	Создайте структуру "Person" с полем "Name" и вложенной структурой
// 	"Address" с полями "Street", "City" и "Country". Добавьте метод
// 	"FullAddress", который возвращает полный адрес как строку.

type Address struct {
	Street  string
	City    string
	Country string
}
type Person struct {
	Name    string
	Address Address
}

func (a Person) FullAddress() string {
	return a.Address.Street + " " + a.Address.City + " " + a.Address.Country
}

func main() {
	p := Person{
		Name: "Joe",
		Address: Address{
			Street:  "Baker",
			City:    "London",
			Country: "Uk",
		},
	}
	fmt.Println(p.FullAddress())
}

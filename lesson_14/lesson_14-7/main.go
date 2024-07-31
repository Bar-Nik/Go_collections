package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 7. Вложенные структуры ( сложное): Создайте структуру Company , которая
// содержит поле Employees , являющееся срезом структур Employee .
// Сериализуйте и десериализуйте экземпляр структуры Company .

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Company struct {
	Company   string     `json:"company"`
	Employees []Employee `json:"employees"`
}

func main() {
	a := Employee{Name: "Jhon", Age: 23}
	b := Employee{Name: "Tom", Age: 43}

	company := Company{Company: "Cola", Employees: []Employee{a, b}}
	fmt.Println("Вывод структуры Company:", company)
	companyJson := mySeri(company)
	fmt.Println("JSON:", companyJson)
	companyC := myDESeri(companyJson)
	fmt.Println("Вывод структуры Company:", companyC)
}

func mySeri(p Company) string {
	bJson, err := json.Marshal(p)
	if err != nil {
		return "Ошибка сериализации"
	}
	return string(bJson)
}

func myDESeri(s string) Company {
	var d Company
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		log.Fatal(err)
	}
	return d

}

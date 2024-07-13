package main

import "fmt"

// 2. Создание структур с разными типами данных: Создайте структуру
// "Person" с полями "Name" (string), "Age" (int) и "IsMarried" (bool). Создайте
// экземпляр этой структуры, заполните поля и выведите их на экран.

type Person struct {
	Name      string
	Age       int
	IsMarried bool
}

func main() {

	p := Person{
		Name:      "Joe Black",
		Age:       30,
		IsMarried: false,
	}
	fmt.Println(p.Name, p.Age, p.IsMarried)
}

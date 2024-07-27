package main

import "fmt"

// 3. Изменение значений полей структуры: Создайте экземпляр структуры
// из предыдущего задания, измените значение одного из полей и выведите
// все поля на экран.

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
	p.Age = 28
	fmt.Println(p.Name, p.Age, p.IsMarried)

}
